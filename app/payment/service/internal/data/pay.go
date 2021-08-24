package data

import (
	"skrshop/app/payment/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type paymentRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewPaymentRepo(data *Data, logger log.Logger) biz.PaymentRepo {
	return &paymentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *paymentRepo) CreateGreeter() error {
	return nil
}

func (r *paymentRepo) UpdateGreeter() error {
	return nil
}
