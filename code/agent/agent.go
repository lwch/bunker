package agent

import (
	"context"

	"github.com/lwch/bunker/code/conf"
	"github.com/lwch/bunker/code/network"
	"github.com/lwch/logging"
	"github.com/lwch/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Run run in agent mode
func Run(cfg *conf.Configure) {
	var conn *grpc.ClientConn
	var err error
	if cfg.UseSSL {
		conn, err = grpc.Dial(cfg.Server)
	} else {
		conn, err = grpc.Dial(cfg.Server, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	runtime.Assert(err)
	defer conn.Close()

	logging.Info("connected on %s", cfg.Server)

	cli := network.NewBunkerClient(conn)

	ctx, cancel := context.WithCancel(context.Background())

	go runShell(ctx, cancel, cli)
	go runVnc(ctx, cancel, cli)

	<-ctx.Done()
}
