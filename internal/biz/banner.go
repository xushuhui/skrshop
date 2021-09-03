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

type BannerUsecase struct {
	repo BannerRepo
	log  *log.Helper
}

func NewBannerUsecase(repo BannerRepo, logger log.Logger) *BannerUsecase {
	return &BannerUsecase{repo: repo, log: log.NewHelper(logger)}
}
