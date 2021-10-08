package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"skrshop/internal/biz"
	"skrshop/internal/data/ent"
	"skrshop/internal/data/ent/theme"
)

var _ biz.ThemeRepo = (*themeRepo)(nil)

type themeRepo struct {
	data *Data
	log  *log.Helper
}

func (r themeRepo) GetThemesByNames(ctx context.Context, names []string) (po []*ent.Theme, err error) {
	po, err = r.data.db.Theme.Query().Where(theme.NameIn(names...)).All(ctx)
	return
}

func (r themeRepo) GetThemeByNameWithSpu(ctx context.Context, name string) (po *ent.Theme, err error) {
	panic("implement me")
}

func NewThemeRepo(data *Data, logger log.Logger) biz.ThemeRepo {
	return &themeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
