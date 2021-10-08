package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"skrshop/internal/data/ent"
)

type BannerRepo interface {
	GetBannerByName(ctx context.Context, name string) (banner *ent.Banner, err error)
	GetBannerById(ctx context.Context, id int64) (banner *ent.Banner, err error)
}

type ShowUsecase struct {
	bannerRepo BannerRepo
	themeRepo  ThemeRepo
	log        *log.Helper
}

func NewShowUsecase(bannerRepo BannerRepo, themeRepo ThemeRepo, logger log.Logger) *ShowUsecase {
	return &ShowUsecase{bannerRepo: bannerRepo, themeRepo: themeRepo,
		log: log.NewHelper(logger)}
}

func (uc *ShowUsecase) GetBannerByName(ctx context.Context, name string) (banner *ent.Banner, err error) {
	banner, err = uc.bannerRepo.GetBannerByName(ctx, name)
	return
}
func (uc *ShowUsecase) GetBannerById(ctx context.Context, id int64) (banner *ent.Banner, err error) {
	banner, err = uc.bannerRepo.GetBannerById(ctx, id)
	return
}
