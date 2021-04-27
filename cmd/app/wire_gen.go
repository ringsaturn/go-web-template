// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/ringsaturn/go-web-template/pkg/config"
	"github.com/ringsaturn/go-web-template/pkg/controllers"
	"github.com/ringsaturn/go-web-template/pkg/dao"
	"github.com/ringsaturn/go-web-template/pkg/server/grpc"
	"github.com/ringsaturn/go-web-template/pkg/server/http"
	"github.com/ringsaturn/go-web-template/pkg/service"
)

// Injectors from wire.go:

func initService(conf *config.Config) (*service.Service, error) {
	daoDao, err := dao.NewDao(conf)
	if err != nil {
		return nil, err
	}
	controller, err := controllers.NewController(daoDao)
	if err != nil {
		return nil, err
	}
	server, err := http.NewServer(conf, controller)
	if err != nil {
		return nil, err
	}
	grpcServer, err := grpc.NewServer()
	if err != nil {
		return nil, err
	}
	serviceService, err := service.NewService(conf, server, grpcServer)
	if err != nil {
		return nil, err
	}
	return serviceService, nil
}
