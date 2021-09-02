// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (


	"skrshop/app/cart/service/internal/biz"
	"skrshop/app/cart/service/internal/conf"
	"skrshop/app/cart/service/internal/data"
	"skrshop/app/cart/service/internal/server"
	"skrshop/app/cart/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)


// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
