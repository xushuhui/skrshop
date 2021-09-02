package service

import (
	"skrshop/app/shop/interface/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	v1 "skrshop/api/cart/v1"
)

// GreeterService is a greeter service.
type CartService struct {
	v1.UnimplementedCartServer

	bc  *biz.BannerUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewCartService(bc *biz.BannerUsecase, logger log.Logger) *CartService {
	return &CartService{bc: bc, log: log.NewHelper(logger)}
}


