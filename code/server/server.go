package server

import (
	"crypto/tls"
	"fmt"
	"net"

	"github.com/lwch/bunker/code/conf"
	"github.com/lwch/bunker/code/network"
	"github.com/lwch/logging"
	"github.com/lwch/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
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
	runtime.Assert(err)
	logging.Info("listen on %d", cfg.Listen)
	svr := grpc.NewServer(options...)
	network.RegisterBunkerServer(svr, &server{})
	runtime.Assert(svr.Serve(l))
}
