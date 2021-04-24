package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ringsaturn/go-web-template/pkg/config"
	"github.com/ringsaturn/go-web-template/pkg/dao"
)

type Server struct {
	HTTPServer *http.Server
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

type Controller interface {
	Hello(*gin.Context)
}

func NewServer(conf *config.Config, dao *dao.Dao, controller Controller) (*Server, error) {
	router := gin.Default()
	router.GET("/", Ping)
	router.GET("/ping", Ping)
	router.POST("/hello", controller.Hello)
	server := &Server{
		HTTPServer: &http.Server{
			Addr:         conf.HTTPServer.Addr,
			Handler:      router,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
	return server, nil
}
