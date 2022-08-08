package service

import (
	"context"

	"github.com/lwch/bunker/code/network"
	"github.com/lwch/logging"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (svr *service) RunShell(ctx context.Context, _ *emptypb.Empty) (*network.RunShellArguments, error) {
	id, err := agentID(ctx)
	if err != nil {
		logging.Error("get agent id for run_shell: %v", err)
		return nil, err
	}
	return <-svr.sh.WaitRun(id), nil
}

func (svr *service) ShellResize(ctx context.Context, _ *emptypb.Empty) (*network.ShellResizeArguments, error) {
	id, err := agentID(ctx)
	if err != nil {
		logging.Error("get agent id for shell_resize: %v", err)
		return nil, err
	}
	return <-svr.sh.WaitResize(id), nil
}

func (svr *service) ShellForward(network.Bunker_ShellForwardServer) error {
	return nil
}
