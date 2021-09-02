package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type BannerRepo interface {
}

type BannerUsecase struct {
	repo BannerRepo
	log  *log.Helper
}

func NewBannerUsecase(repo BannerRepo, logger log.Logger) *BannerUsecase {
	return &BannerUsecase{repo: repo, log: log.NewHelper(logger)}
}
