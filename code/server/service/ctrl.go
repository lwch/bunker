package service

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/lwch/bunker/code/network"
	"github.com/lwch/logging"
	"github.com/lwch/runtime"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (svc *service) Connect(ctx context.Context, args *network.ConnectArgs) (*network.ConnectResponse, error) {
	var ret network.ConnectResponse
	if len(args.GetId()) == 0 {
		id, err := runtime.UUID(16, "0123456789abcdef")
		if err != nil {
			logging.Error("generate agent_id: %v", err)
			return nil, err
		}
		ret.Id = fmt.Sprintf("agent-%s-%04d-%s",
			time.Now().Format("20060102"),
			atomic.AddUint64(&svc.idPrefix, 1), id)
	} else {
		ret.Id = args.GetId()
	}
	ret.Version = svc.version
	return &ret, nil
}

func (svc *service) KeepAlive(ctx context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	id, err := agentID(ctx)
	if err != nil {
		logging.Error("keepalive: %v", err)
		return nil, err
	}
	for _, h := range svc.handler {
		go h.OnKeepalive(id)
	}
	return &emptyResp, nil
}
