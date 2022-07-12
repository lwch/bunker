package agent

import (
	"context"

	"github.com/lwch/bunker/code/conf"
	"github.com/lwch/bunker/code/network"
	"github.com/lwch/bunker/code/utils"
	"github.com/lwch/logging"
	"github.com/lwch/runtime"
	"google.golang.org/protobuf/types/known/emptypb"
)

func runShell(ctx context.Context, cancel context.CancelFunc, cli network.BunkerClient, cfg *conf.Configure) {
	defer utils.Recover("shell")
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			ctx = cfg.SecretContext(ctx)
			args, err := cli.RunShell(ctx, &emptypb.Empty{})
			runtime.Assert(err)
			logging.Info("args: %v", args)
		}
	}
}
