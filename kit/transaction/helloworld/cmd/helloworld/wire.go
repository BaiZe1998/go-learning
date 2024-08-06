//go:build wireinject
// +build wireinject

//go:generate wire

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"helloworld/internal/biz"
	"helloworld/internal/conf"
	"helloworld/internal/data"
	"helloworld/internal/server"
	"helloworld/internal/service"
	"helloworld/pkg/db"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(db.NewDBClient, server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
