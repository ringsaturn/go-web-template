// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ringsaturn/go-web-template/pkg/config"
	"github.com/ringsaturn/go-web-template/pkg/controllers"
	"github.com/ringsaturn/go-web-template/pkg/dao"
	pkggrpc "github.com/ringsaturn/go-web-template/pkg/server/grpc"
	pkghttp "github.com/ringsaturn/go-web-template/pkg/server/http"
	"github.com/ringsaturn/go-web-template/pkg/service"
)

func initService(conf *config.Config) (*service.Service, error) {
	panic(wire.Build(
		dao.ProviderSet,
		pkghttp.ProviderSet,
		pkggrpc.ProviderSet,
		service.ProviderSet,
		controllers.ProviderSet,
	))
}
