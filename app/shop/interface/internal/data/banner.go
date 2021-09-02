package data

import (
	"context"
	"skrshop/app/shop/interface/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type bannerRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewBannerRepo(data *Data, logger log.Logger) biz.BannerRepo {
	return &bannerRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u bannerRepo) DeleteCart(ctx context.Context, uid int64) error {
	panic("implement me")
}
