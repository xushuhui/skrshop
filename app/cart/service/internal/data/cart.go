package data

import (
	"context"
	"skrshop/app/cart/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type cartRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewCartRepo(data *Data, logger log.Logger) biz.CartRepo {

	return &cartRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u cartRepo)GetCart(ctx context.Context, uid int64) (*biz.Cart, error) {
	panic("implement me")
}

func (u cartRepo) SaveCart(ctx context.Context, c *biz.Cart) error{
	panic("implement me")
}

func (u cartRepo) DeleteCart(ctx context.Context, uid int64)  error {
	panic("implement me")
}


