package controllers

import (
	"github.com/google/wire"
	pkghttp "github.com/ringsaturn/go-web-template/pkg/server/http"
)

var ProviderSet = wire.NewSet(
	NewController,
	wire.Bind(new(pkghttp.Controller), new(*Controller)),
)
