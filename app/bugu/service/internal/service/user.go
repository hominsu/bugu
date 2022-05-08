package service

import (
	"context"

	"github.com/google/uuid"

	buguV1 "bugu/api/bugu/service/v1"
)

func (s *BuguService) Register(ctx context.Context, in *buguV1.RegisterRequest) (*buguV1.RegisterReply, error) {
	user := in.GetUser()

	dto, err := s.uu.Register(ctx, user.Email, user.Username, user.Password)
	if err != nil {
		return nil, err
	}

	return &buguV1.RegisterReply{User: &buguV1.UserStruct{
		Id:       dto.ID.String(),
		Email:    dto.Email,
		Username: dto.Username,
	}}, nil
}

func (s *BuguService) Login(ctx context.Context, in *buguV1.LoginRequest) (*buguV1.LoginReply, error) {
	user := in.GetUser()

	dto, err := s.uu.Login(ctx, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	return &buguV1.LoginReply{
		User: &buguV1.UserStruct{
			Id:       dto.ID.String(),
			Email:    dto.Email,
			Username: dto.Username,
		},
		Token: dto.Token,
	}, nil
}

func (s *BuguService) GetCurrentUser(ctx context.Context, in *buguV1.GetCurrentUserRequest) (*buguV1.GetCurrentUserReply, error) {
	u, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, buguV1.ErrorUuidParseFailed("parse userid failed, id: %s", in.Id)
	}

	dto, err := s.uu.GetUserByID(ctx, u)
	if err != nil {
		return nil, err
	}

	return &buguV1.GetCurrentUserReply{User: &buguV1.UserStruct{
		Id:       dto.ID.String(),
		Email:    dto.Email,
		Username: dto.Username,
	}}, nil
}

func (s *BuguService) UpdateUser(ctx context.Context, in *buguV1.UpdateUserRequest) (*buguV1.UpdateUserReply, error) {
	user := in.GetUser()

	dto, err := s.uu.UpdateUser(ctx, user.Id, user.Email, user.Username, user.Password)
	if err != nil {
		return nil, err
	}

	return &buguV1.UpdateUserReply{User: &buguV1.UserStruct{
		Id:       dto.ID.String(),
		Email:    dto.Email,
		Username: dto.Username,
	}}, nil
}
