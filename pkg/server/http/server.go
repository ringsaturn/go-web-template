package http

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ringsaturn/go-web-template/pkg/config"
)

type Server struct {
	httpServer *http.Server
}

type Controller interface {
	Ping(*gin.Context)
	Hello(*gin.Context)
}

func NewServer(conf *config.Config, controller Controller) (*Server, error) {
	router := gin.Default()
	router.GET("/", controller.Ping)
	router.GET("/ping", controller.Ping)
	router.POST("/hello", controller.Hello)
	server := &Server{
		httpServer: &http.Server{
			Addr:         conf.HTTPServer.Addr,
			Handler:      router,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
	return server, nil
}

func (srv *Server) Start(ctx context.Context) error {
	return srv.httpServer.ListenAndServe()
}

func (srv *Server) Shutdown(ctx context.Context) error {
	return srv.httpServer.Shutdown(ctx)
}
