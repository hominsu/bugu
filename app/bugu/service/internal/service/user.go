/*
 * MIT License
 *
 * Copyright (c) 2022. HominSu
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */

package service

import (
	"context"

	"github.com/google/uuid"

	buguV1 "github.com/hominsu/bugu/api/bugu/service/v1"
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
