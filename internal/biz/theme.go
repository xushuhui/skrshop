package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"skrshop/internal/data/ent"
)

type ThemeRepo interface {
	GetThemesByNames(ctx context.Context, names []string) (themes []*ent.Theme, err error)
	GetThemeByNameWithSpu(ctx context.Context, name string) (theme *ent.Theme, err error)
}

//todo 预留？
type ThemeUsecase struct {
	repo ThemeRepo
	log  *log.Helper
}

func NewThemeUsecase(repo ThemeRepo, logger log.Logger) *ThemeUsecase {
	return &ThemeUsecase{repo: repo, log: log.NewHelper(logger)}
}
func (uc *ShowUsecase) GetThemesByNames(ctx context.Context, names []string) (themes []*ent.Theme, err error) {
	themes, err = uc.themeRepo.GetThemesByNames(ctx, names)
	return
}
func (uc *ShowUsecase) GetThemeByNameWithSpu(ctx context.Context, name string) (theme *ent.Theme, err error) {
	theme, err = uc.themeRepo.GetThemeByNameWithSpu(ctx, name)
	return
}
