package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "skrshop/api/user/v1"
	"skrshop/app/user/service/internal/biz"
)

// GreeterService is a greeter service.
type UserService struct {
	v1.UnimplementedUserServer

	uc  *biz.UserUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

// SayHello implements helloworld.GreeterServer
func (s *UserService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	s.log.WithContext(ctx).Infof("SayHello Received: %v", in.GetName())

	if in.GetName() == "error" {
		return nil, errors.NotFound(v1.Error_USER_NOT_FOUND.String(), in.GetName())
	}
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}
