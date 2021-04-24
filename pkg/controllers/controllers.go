package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/ringsaturn/go-web-template/api"
	"github.com/ringsaturn/go-web-template/pkg/dao"
)

type Controller struct {
	d *dao.Dao
}

func NewController(d *dao.Dao) (*Controller, error) {
	controllers := &Controller{
		d: d,
	}
	return controllers, nil
}

func (controllers *Controller) Hello(c *gin.Context) {
	resp := &api.HelloResp{}
	resp.Status = false

	data := &api.HelloResp{}
	err := jsonpb.Unmarshal(c.Request.Body, data)

	if err != nil {
		resp.Code = api.HelloResponseCode(api.HelloResponseCode_UNMARSHAL_FAILED.Number())
		resp.Msg = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Status = true
	resp.Code = api.HelloResponseCode(api.HelloResponseCode_OK.Number())
	c.JSON(http.StatusCreated, resp)
}
