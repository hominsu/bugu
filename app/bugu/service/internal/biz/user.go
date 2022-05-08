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
