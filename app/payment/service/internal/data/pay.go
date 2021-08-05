package data

import (
	"context"
	"skrshop/app/payment/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type paymentRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.PaymentRepo{
	return &paymentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *paymentRepo) CreateGreeter(ctx context.Context) error {
	return nil
}

func (r *paymentRepo) UpdateGreeter(ctx context.Context) error {
	return nil
}
