package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type Payment struct {
	Hello string
}

type PaymentRepo interface {
	
}

type PaymentUsecase struct {
	repo PaymentRepo
	log  *log.Helper
}

func NewPaymentUsecase(repo PaymentRepo, logger log.Logger) *PaymentUsecase {
	return &PaymentUsecase{repo: repo, log: log.NewHelper(logger)}
}

