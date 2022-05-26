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

package biz

import (
	"context"

	buguV1 "bugu/api/bugu/service/v1"
	"bugu/app/bugu/service/internal/conf"
	"bugu/app/bugu/service/internal/pkg/middleware/auth"
	"bugu/pkg"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

// User is the model entity for the User schema.
type User struct {
	ID           uuid.UUID `json:"id,omitempty"`
	Email        string    `json:"email,omitempty"`
	Username     string    `json:"username,omitempty"`
	PasswordHash string    `json:"password_hash,omitempty"`
	Token        string    `json:"token,omitempty"`
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*User, error)
	FineUserByEmail(ctx context.Context, email string) (*User, error)
}

type UserUsecase struct {
	repo UserRepo
	jc   *conf.Jwt

	log *log.Helper
}

func NewUserUsecase(repo UserRepo, jc *conf.Jwt, logger log.Logger) *UserUsecase {
	return &UserUsecase{
		repo: repo,
		jc:   jc,
		log:  log.NewHelper(logger),
	}
}

func (uc *UserUsecase) Login(ctx context.Context, email, password string) (*User, error) {
	dto, err := uc.repo.FineUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if !pkg.VerifyPassword(dto.PasswordHash, password) {
		return nil, buguV1.ErrorLoginFailed("Verify password failed")
	}

	token, err := auth.GenerateToken(uc.jc.Secret, dto.ID.String())
	if err != nil {
		return nil, err
	}
	dto.Token = token

	return dto, err
}

func (uc *UserUsecase) Register(ctx context.Context, email, username, password string) (*User, error) {
	ph, err := pkg.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &User{
		Email:        email,
		Username:     username,
		PasswordHash: ph,
	}

	a, err := uc.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, id, email, username, password string) (*User, error) {
	u, err := uuid.Parse(id)
	if err != nil {
		return nil, buguV1.ErrorUuidParseFailed("parse userid failed, id: %s", id)
	}

	user := &User{
		ID:       u,
		Email:    email,
		Username: username,
	}

	if password != "" {
		ph, err := pkg.HashPassword(password)
		if err != nil {
			return nil, err
		}
		user.PasswordHash = ph
	}

	return uc.repo.UpdateUser(ctx, user)
}

func (uc *UserUsecase) GetUserByID(ctx context.Context, id uuid.UUID) (*User, error) {
	return uc.repo.GetUserByID(ctx, id)
}
