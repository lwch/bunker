package service

import (
	"context"
	"crypto/tls"
	"embed"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lwch/bunker/code/conf"
	"github.com/lwch/bunker/code/network"
	"github.com/lwch/bunker/code/server/shell"
	"github.com/lwch/bunker/code/utils"
	"github.com/lwch/logging"
	"github.com/lwch/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/types/known/emptypb"
)

var emptyResp emptypb.Empty

type handler interface {
	ApiFuncs() []gin.RouteInfo
	OnConnect(string)
	OnDisconnect(string)
	OnKeepalive(string)
}

type service struct {
	network.UnimplementedBunkerServer
	cfg      *conf.Configure
	idPrefix uint64
	version  string

	// handlers
	handler []handler
	sh      *shell.Handler
}

// New create service
func New(cfg *conf.Configure, version string) *service {
	sh := shell.New()
	svc := &service{
		cfg:     cfg,
		version: version,
		sh:      sh,
	}
	svc.handler = append(svc.handler, sh)
	return svc
}

// Run run in server mode
func (svc *service) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go svc.grpcServe(ctx, cancel)
	go svc.httpServe(ctx, cancel)

	<-ctx.Done()
}

func (svc *service) grpcServe(ctx context.Context, cancel context.CancelFunc) {
	defer utils.Recover("grpc_serve")
	defer cancel()

	var l net.Listener
	var err error
	var options []grpc.ServerOption
	if len(svc.cfg.TLSCrt) > 0 && len(svc.cfg.TLSKey) > 0 {
		cert, err := tls.LoadX509KeyPair(svc.cfg.TLSCrt, svc.cfg.TLSKey)
		runtime.Assert(err)
		options = append(options, grpc.Creds(credentials.NewServerTLSFromCert(&cert)))
	}
	l, err = net.Listen("tcp", fmt.Sprintf(":%d", svc.cfg.GrpcListen))
	options = append(options, grpc.UnaryInterceptor(svc.verify))
	runtime.Assert(err)
	logging.Info("grpc listen on %d", svc.cfg.GrpcListen)
	s := grpc.NewServer(options...)
	network.RegisterBunkerServer(s, svc)
	runtime.Assert(s.Serve(l))
}

//go:generate cp -r ../../../frontend/dist dist
//go:embed dist
var html embed.FS

func (svc *service) httpServe(ctx context.Context, cancel context.CancelFunc) {
	defer utils.Recover("http_serve")
	defer cancel()

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		logging.Error("http handle error: %v", err)
		var e error
		switch er := err.(type) {
		case error:
			e = er
		default:
			e = fmt.Errorf("%v", err)
		}
		c.AbortWithError(http.StatusInternalServerError, e)
	}))

	dist, err := fs.Sub(html, "dist")
	runtime.Assert(err)
	router.StaticFS("/", http.FS(dist))

	api := router.Group("/api")
	reg := func(h handler) {
		for _, info := range h.ApiFuncs() {
			api.Handle(info.Method, info.Path, info.HandlerFunc)
		}
	}
	for _, h := range svc.handler {
		reg(h)
	}

	logging.Info("http listen on %d", svc.cfg.HttpListen)
	if len(svc.cfg.TLSCrt) > 0 && len(svc.cfg.TLSKey) > 0 {
		err = router.RunTLS(fmt.Sprintf(":%d", svc.cfg.HttpListen),
			svc.cfg.TLSCrt, svc.cfg.TLSKey)
	} else {
		err = router.Run(fmt.Sprintf(":%d", svc.cfg.HttpListen))
	}
	runtime.Assert(err)
}

func (svc *service) verify(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	err := svc.cfg.SecretVerify(ctx)
	if err != nil {
		time.Sleep(time.Second)
		return nil, err
	}
	return handler(ctx, req)
}
