//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"applet-server/internal/biz"
	"applet-server/internal/conf"
	"applet-server/internal/data"
	"applet-server/internal/pkg/log"
	"applet-server/internal/server"
	"applet-server/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.App, *conf.Data, *conf.Log) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, log.ProviderSet, service.ProviderSet, biz.ProviderSet, data.ProviderSet, newApp))
}
