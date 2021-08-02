package service

import (
	"skrshop/app/cart/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	v1 "skrshop/api/cart/v1"
)

// GreeterService is a greeter service.
type CartService struct {
	v1.UnimplementedCartServer

	uc  *biz.CartUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewCartService(uc *biz.CartUsecase, logger log.Logger) *CartService {
	return &CartService{uc: uc, log: log.NewHelper(logger)}
}

// SayHello implements helloworld.GreeterServer
