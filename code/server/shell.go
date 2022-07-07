package server

import (
	"context"

	"github.com/lwch/bunker/code/network"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (svr *server) RunShell(context.Context, *emptypb.Empty) (*network.RunShellArguments, error) {
	return nil, nil
}

func (svr *server) ShellResize(context.Context, *emptypb.Empty) (*network.ShellResizeArguments, error) {
	return nil, nil
}

func (svr *server) ShellForward(network.Bunker_ShellForwardServer) error {
	return nil
}
