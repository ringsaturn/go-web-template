// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ringsaturn/go-web-template/pkg/config"
	"github.com/ringsaturn/go-web-template/pkg/controllers"
	"github.com/ringsaturn/go-web-template/pkg/dao"
	"github.com/ringsaturn/go-web-template/pkg/server"
	"github.com/ringsaturn/go-web-template/pkg/service"
)

func initService(conf *config.Config) (*service.Service, error) {
	panic(wire.Build(dao.ProviderSet, server.ProviderSet, service.ProviderSet, controllers.ProviderSet))
}
