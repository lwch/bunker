package shell

import (
	"context"

	"github.com/lwch/bunker/code/conf"
	"github.com/lwch/bunker/code/network"
	"github.com/lwch/logging"
	"github.com/lwch/runtime"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Handler shell handler
type Handler struct {
	cfg *conf.Configure
	cli network.BunkerClient
}

// New create shell handler
func New(cfg *conf.Configure, cli network.BunkerClient) *Handler {
	return &Handler{cfg: cfg, cli: cli}
}

// Run run shell handler
func (h *Handler) Run(ctx context.Context) error {
	args, err := h.cli.RunShell(ctx, &emptypb.Empty{})
	runtime.Assert(err)
	logging.Info("args: %v", args)
	return nil
}

// Resize resize shell handler
func (h *Handler) Resize(ctx context.Context) error {
	args, err := h.cli.ShellResize(ctx, &emptypb.Empty{})
	runtime.Assert(err)
	logging.Info("args: %v", args)
	return nil
}
