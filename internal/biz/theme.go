package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"skrshop/internal/data/ent"
)

type ThemeRepo interface {
	GetThemesByNames(ctx context.Context, names []string) (themes *ent.Themes, err error)
	GetThemeByNameWithSpu(ctx context.Context, name string) (theme *ent.Theme, err error)
}

type ThemeUsecase struct {
	repo ThemeRepo
	log  *log.Helper
}

func NewThemeUsecase(repo ThemeRepo, logger log.Logger) *ThemeUsecase {
	return &ThemeUsecase{repo: repo, log: log.NewHelper(logger)}
}
