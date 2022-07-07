package agent

import (
	"context"
	"time"

	"github.com/lwch/bunker/code/network"
	"github.com/lwch/bunker/code/utils"
)

func runVnc(ctx context.Context, cancel context.CancelFunc, cli network.BunkerClient) {
	defer utils.Recover("vnc")
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// TODO
			time.Sleep(time.Second)
		}
	}
}
