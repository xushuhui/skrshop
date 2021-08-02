//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"skrshop/app/user/service/internal/biz"
	"skrshop/app/user/service/internal/conf"
	"skrshop/app/user/service/internal/data"
	"skrshop/app/user/service/internal/server"
	"skrshop/app/user/service/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	repo := data.NewUserRepo(dataData, logger)
	usecase := biz.NewUserUsecase(repo, logger)
	userService := service.NewUserService(usecase, logger)
	httpServer := server.NewHTTPServer(confServer, userService, logger)
	grpcServer := server.NewGRPCServer(confServer, userService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
