package service

import (
	"context"

	buguV1 "bugu/api/bugu/service/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *BuguService) Detect(ctx context.Context, in *buguV1.DetectRequest) (*buguV1.DetectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detect not implemented")
}
