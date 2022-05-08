//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"bugu/app/bugu/service/internal/biz"
	"bugu/app/bugu/service/internal/conf"
	"bugu/app/bugu/service/internal/data"
	"bugu/app/bugu/service/internal/server"
	"bugu/app/bugu/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Registry, *conf.Server, *conf.Data, *conf.Jwt, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
