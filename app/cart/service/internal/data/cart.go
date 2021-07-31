package data

import (
	"skrshop/internal/biz"
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

func (r *cartRepo) CreateGreeter(ctx context.Context, g *biz.Greeter) error {
	return nil
}

func (r *cartRepo) UpdateGreeter(ctx context.Context, g *biz.Greeter) error {
	return nil
}
