package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	UserId int64
}
type UserRepo interface {
	GetUser(ctx context.Context, uid int64) (*User, error)
	SaveUser(ctx context.Context, c *User) error
	DeleteUser(ctx context.Context, uid int64) error
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) GetUser(ctx context.Context, uid int64) (*User, error) {
	return uc.repo.GetUser(ctx, uid)
}

func (uc *UserUsecase) DeleteUser(ctx context.Context, uid int64) error {
	return uc.repo.DeleteUser(ctx, uid)
}
