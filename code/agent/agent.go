package agent

import (
	"context"
	"os"

	"github.com/lwch/bunker/code/conf"
	"github.com/lwch/bunker/code/network"
	"github.com/lwch/logging"
	"github.com/lwch/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Run run in agent mode
func Run(cfg *conf.Configure) {
	var options []grpc.DialOption
	if !cfg.UseSSL {
		options = append(options, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	options = append(options, grpc.WithUnaryInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx = cfg.SecretContext(ctx)
		return invoker(ctx, method, req, reply, cc, opts...)
	}))

	conn, err := grpc.Dial(cfg.Server, options...)
	runtime.Assert(err)
	defer conn.Close()

	logging.Info("connected on %s", cfg.Server)

	cli := network.NewBunkerClient(conn)

	ctx, cancel := context.WithCancel(context.Background())

	go runShell(ctx, cancel, cli, cfg)
	go runVnc(ctx, cancel, cli)

	<-ctx.Done()
	os.Exit(1)
}
