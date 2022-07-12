package vnc

import (
	"context"

	"github.com/lwch/bunker/code/conf"
	"github.com/lwch/bunker/code/network"
)

// Handler vnc handler
type Handler struct {
	cfg *conf.Configure
	cli network.BunkerClient
}

// New create vnc handler
func New(cfg *conf.Configure, cli network.BunkerClient) *Handler {
	return &Handler{cfg: cfg, cli: cli}
}

// Run run vnc handler
func (h *Handler) Run(ctx context.Context) error {
	return nil
}
