package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func agentID(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.InvalidArgument, "Retrieving metadata is failed")
	}
	id := md.Get("id")
	if len(id) == 0 {
		return "", status.Errorf(codes.Unauthenticated, "Missing agent id")
	}
	return id[0], nil
}
