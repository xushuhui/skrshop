package data

import (
	"context"
	"skrshop/app/shop/interface/internal/biz"
	"skrshop/app/shop/interface/internal/data/ent"

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

func (u bannerRepo) GetBannerByName(ctx context.Context, name string) (banner *ent.Banner, err error) {
	panic("implement me")
}
func (u bannerRepo) GetBannerById(ctx context.Context, id int64) (banner *ent.Banner, err error) {
	panic("implement me")
}
