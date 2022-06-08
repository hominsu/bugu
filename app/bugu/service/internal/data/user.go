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

package data

import (
	"context"
	"time"

	buguV1 "github.com/hominsu/bugu/api/bugu/service/v1"
	"github.com/hominsu/bugu/app/bugu/service/internal/biz"
	"github.com/hominsu/bugu/app/bugu/service/internal/data/ent"
	"github.com/hominsu/bugu/app/bugu/service/internal/data/ent/user"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user")),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return nil, buguV1.ErrorUuidGenerateFailed("create follow uuid failed, err: %v", err)
	}

	po, err := r.data.db.User.Create().
		SetID(u).
		SetEmail(user.Email).
		SetUsername(user.Username).
		SetPasswordHash(user.PasswordHash).
		Save(ctx)
	if err != nil && ent.IsConstraintError(err) {
		return nil, buguV1.ErrorCreateConflict("create conflict, err: %v", err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, buguV1.ErrorUnknownError("unknown err: %v", err)
	}

	return &biz.User{
		ID:           po.ID,
		Email:        po.Email,
		Username:     po.Username,
		PasswordHash: po.PasswordHash,
	}, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	po, err := r.data.db.User.UpdateOneID(user.ID).
		SetUsername(user.Username).
		SetPasswordHash(user.PasswordHash).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil && ent.IsConstraintError(err) {
		return nil, buguV1.ErrorCreateConflict("update conflict, err: %v", err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, buguV1.ErrorUnknownError("unknown err: %v", err)
	}

	return &biz.User{
		ID:           po.ID,
		Email:        po.Email,
		Username:     po.Username,
		PasswordHash: po.PasswordHash,
	}, nil
}

func (r *userRepo) GetUserByID(ctx context.Context, id uuid.UUID) (*biz.User, error) {
	po, err := r.data.db.User.Get(ctx, id)
	if err != nil && ent.IsNotFound(err) {
		return nil, buguV1.ErrorNotFoundError("find user id: %s not found, err: %v", id.String(), err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, buguV1.ErrorUnknownError("unknown err: %v", err)
	}

	return &biz.User{
		ID:           po.ID,
		Email:        po.Email,
		Username:     po.Username,
		PasswordHash: po.PasswordHash,
	}, nil
}

func (r *userRepo) FineUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	target, err := r.data.db.User.Query().
		Where(user.EmailEQ(email)).
		Only(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, buguV1.ErrorNotFoundError("find user email: %s not found, err: %v", email, err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, buguV1.ErrorUnknownError("unknown err: %v", err)
	}

	return &biz.User{
		ID:           target.ID,
		Email:        target.Email,
		Username:     target.Username,
		PasswordHash: target.PasswordHash,
	}, nil
}
