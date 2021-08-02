package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"skrshop/app/user/service/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u userRepo) GetUser(ctx context.Context, uid int64) (*biz.User, error) {
	panic("implement me")
}

func (u userRepo) SaveUser(ctx context.Context, c *biz.User) error {
	panic("implement me")
}

func (u userRepo) DeleteUser(ctx context.Context, uid int64) error {
	panic("implement me")
}
