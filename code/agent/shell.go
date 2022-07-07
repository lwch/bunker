package agent

import (
	"context"

	"github.com/lwch/bunker/code/network"
	"github.com/lwch/bunker/code/utils"
)

func runShell(ctx context.Context, cancel context.CancelFunc, cli network.BunkerClient) {
	defer utils.Recover("shell")
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			cli.RunShell(ctx, nil)
		}
	}
}
