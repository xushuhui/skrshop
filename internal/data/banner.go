package data

import (
	"context"
	"skrshop/internal/biz"
	"skrshop/internal/data/ent"
	"skrshop/internal/data/ent/banner"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.BannerRepo = (*bannerRepo)(nil)

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

func (r bannerRepo) GetBannerByName(ctx context.Context, name string) (model *ent.Banner, err error) {
	model, err = r.data.db.Banner.Query().Where(banner.Name(name)).First(ctx)
	return
}
func (r bannerRepo) GetBannerById(ctx context.Context, id int64) (model *ent.Banner, err error) {
	model, err = r.data.db.Banner.Query().Where(banner.ID(id)).First(ctx)
	return

}
