package controllers

import (
	"github.com/google/wire"
	"github.com/ringsaturn/go-web-template/pkg/server"
)

var ProviderSet = wire.NewSet(
	NewController,
	wire.Bind(new(server.Controller), new(*Controller)),
)
