package main

import (
	"context"
	"os"

	"github.com/lwch/bunker/code/agent/shell"
	"github.com/lwch/bunker/code/agent/vnc"
	"github.com/lwch/bunker/code/conf"
	"github.com/lwch/bunker/code/network"
	"github.com/lwch/bunker/code/utils"
	"github.com/lwch/logging"
	"github.com/lwch/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type agent struct {
	cfg *conf.Configure
}

func newAgent(cfg *conf.Configure) *agent {
	return &agent{cfg: cfg}
}

// Run run in agent mode
func (a *agent) Run() {
	var options []grpc.DialOption
	if !a.cfg.UseSSL {
		options = append(options, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	options = append(options, grpc.WithUnaryInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx = a.cfg.BuildHeader(ctx)
		return invoker(ctx, method, req, reply, cc, opts...)
	}))

	conn, err := grpc.Dial(a.cfg.Server, options...)
	runtime.Assert(err)
	defer conn.Close()

	logging.Info("connected on %s", a.cfg.Server)

	cli := network.NewBunkerClient(conn)

	ctx, cancel := context.WithCancel(context.Background())

	loop := func(name string, cb func(context.Context) error) {
		defer utils.Recover(name)
		defer cancel()
		for {
			select {
			case <-ctx.Done():
			default:
				err := cb(ctx)
				if err != nil {
					logging.Error("call %s failed: %s", err)
					return
				}
			}
		}
	}

	sh := shell.New(a.cfg, cli)
	vn := vnc.New(a.cfg, cli)

	go loop("run_shell", sh.Run)
	go loop("shell_resize", sh.Resize)
	go loop("run_vnc", vn.Run)

	<-ctx.Done()
	os.Exit(1)
}
