package service

import (


	v1 "skrshop/api/payment/v1"
	"skrshop/app/payment/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// GreeterService is a greeter service.
type PaymentService struct {
	v1.UnimplementedPaymentServer

	uc  *biz.PaymentUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewPaymentService(uc *biz.PaymentUsecase, logger log.Logger) *PaymentService {
	
	return &PaymentService{uc: uc, log: log.NewHelper(logger)}
}

