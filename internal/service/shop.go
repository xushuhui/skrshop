package service

import (
	"skrshop/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	v1 "skrshop/api/shop/v1"
)

// GreeterService is a greeter service.
type ShopService struct {
	v1.UnimplementedShopServer

	bc *biz.ShowUsecase

	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewShopService(bc *biz.ShowUsecase, logger log.Logger) *ShopService {
	return &ShopService{bc: bc, log: log.NewHelper(logger)}
}
