package main

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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type server struct {
	network.UnimplementedBunkerServer
	cfg *conf.Configure
	sh  *shell.Handler
}

func newServer(cfg *conf.Configure) *server {
	return &server{
		cfg: cfg,
		sh:  shell.New(),
	}
}

// Run run in server mode
func (svr *server) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go svr.grpcServe(ctx, cancel)
	go svr.httpServe(ctx, cancel)

	<-ctx.Done()
}

func (svr *server) grpcServe(ctx context.Context, cancel context.CancelFunc) {
	defer utils.Recover("grpc_serve")
	defer cancel()

	var l net.Listener
	var err error
	var options []grpc.ServerOption
	if len(svr.cfg.TLSCrt) > 0 && len(svr.cfg.TLSKey) > 0 {
		cert, err := tls.LoadX509KeyPair(svr.cfg.TLSCrt, svr.cfg.TLSKey)
		runtime.Assert(err)
		options = append(options, grpc.Creds(credentials.NewServerTLSFromCert(&cert)))
	}
	l, err = net.Listen("tcp", fmt.Sprintf(":%d", svr.cfg.GrpcListen))
	options = append(options, grpc.UnaryInterceptor(svr.verify))
	runtime.Assert(err)
	logging.Info("grpc listen on %d", svr.cfg.GrpcListen)
	s := grpc.NewServer(options...)
	network.RegisterBunkerServer(s, svr)
	runtime.Assert(s.Serve(l))
}

//go:generate cp -r ../../frontend/dist dist
//go:embed dist
var html embed.FS

func (svr *server) httpServe(ctx context.Context, cancel context.CancelFunc) {
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

	type handler interface {
		ApiFuncs() []gin.RouteInfo
	}

	api := router.Group("/api")
	reg := func(h handler) {
		for _, info := range h.ApiFuncs() {
			api.Handle(info.Method, info.Path, info.HandlerFunc)
		}
	}
	reg(svr.sh)

	logging.Info("http listen on %d", svr.cfg.HttpListen)
	if len(svr.cfg.TLSCrt) > 0 && len(svr.cfg.TLSKey) > 0 {
		err = router.RunTLS(fmt.Sprintf(":%d", svr.cfg.HttpListen),
			svr.cfg.TLSCrt, svr.cfg.TLSKey)
	} else {
		err = router.Run(fmt.Sprintf(":%d", svr.cfg.HttpListen))
	}
	runtime.Assert(err)
}

func (svr *server) verify(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	err := svr.cfg.SecretVerify(ctx)
	if err != nil {
		time.Sleep(time.Second)
		return nil, err
	}
	return handler(ctx, req)
}

func agentID(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.InvalidArgument, "Retrieving metadata is failed")
	}
	id := md.Get("id")
	if len(id) == 0 {
		return "", status.Errorf(codes.Unauthenticated, "Missing agent id")
	}
	return id[0], nil
}
