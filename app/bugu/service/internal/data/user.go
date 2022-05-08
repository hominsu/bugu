package data

import (
	"context"
	"time"

	v1 "bugu/api/bugu/service/v1"
	"bugu/app/bugu/service/internal/biz"
	"bugu/app/bugu/service/internal/data/ent"
	"bugu/app/bugu/service/internal/data/ent/user"

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
		return nil, v1.ErrorUuidGenerateFailed("create follow uuid failed, err: %v", err)
	}

	po, err := r.data.db.User.Create().
		SetID(u).
		SetEmail(user.Email).
		SetUsername(user.Username).
		SetPasswordHash(user.PasswordHash).
		Save(ctx)
	if err != nil && ent.IsConstraintError(err) {
		return nil, v1.ErrorCreateConflict("create conflict, err: %v", err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, v1.ErrorUnknownError("unknown err: %v", err)
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
		return nil, v1.ErrorCreateConflict("update conflict, err: %v", err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, v1.ErrorUnknownError("unknown err: %v", err)
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
		return nil, v1.ErrorNotFoundError("find user id: %s not found, err: %v", id.String(), err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, v1.ErrorUnknownError("unknown err: %v", err)
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
		return nil, v1.ErrorNotFoundError("find user email: %s not found, err: %v", email, err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, v1.ErrorUnknownError("unknown err: %v", err)
	}

	return &biz.User{
		ID:           target.ID,
		Email:        target.Email,
		Username:     target.Username,
		PasswordHash: target.PasswordHash,
	}, nil
}
