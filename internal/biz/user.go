package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"skrshop/internal/data/ent"
)

type UserRepo interface {
	GetUserByID(ctx context.Context, id int64) (user *ent.User, err error)
}
type UserUsecase struct {
	repo ThemeRepo
	log  *log.Helper
}

func NewUserUsecase(repo ThemeRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}
