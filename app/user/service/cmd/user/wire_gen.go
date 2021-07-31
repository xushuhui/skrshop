

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
	dataData,cleanup, err := data.NewData(confData)
	if err != nil {
		return nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	greeterService := service.NewGreeterService(greeterUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, greeterService)
	grpcServer := server.NewGRPCServer(confServer, greeterService)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
