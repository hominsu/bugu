package service

import (
	"context"

	buguV1 "bugu/api/bugu/service/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *BuguService) Confusion(ctx context.Context, in *buguV1.ConfusionRequest) (*buguV1.ConfusionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Confusion not implemented")
}
