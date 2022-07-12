package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"time"

	"github.com/lwch/bunker/code/conf"
	"github.com/lwch/bunker/code/network"
	"github.com/lwch/logging"
	"github.com/lwch/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	cfg *conf.Configure
	network.UnimplementedBunkerServer
}

// Run run in server mode
func Run(cfg *conf.Configure) {
	var l net.Listener
	var err error
	var options []grpc.ServerOption
	if len(cfg.TLSCrt) > 0 && len(cfg.TLSKey) > 0 {
		cert, err := tls.LoadX509KeyPair(cfg.TLSCrt, cfg.TLSKey)
		runtime.Assert(err)
		options = append(options, grpc.Creds(credentials.NewServerTLSFromCert(&cert)))
	} else {
		l, err = net.Listen("tcp", fmt.Sprintf(":%d", cfg.Listen))
	}
	s := &server{cfg: cfg}
	options = append(options, grpc.UnaryInterceptor(s.verify))
	runtime.Assert(err)
	logging.Info("listen on %d", cfg.Listen)
	svr := grpc.NewServer(options...)
	network.RegisterBunkerServer(svr, &server{})
	runtime.Assert(svr.Serve(l))
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
