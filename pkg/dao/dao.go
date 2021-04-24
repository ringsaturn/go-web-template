package dao

import (
	"github.com/ringsaturn/go-web-template/pkg/config"
)

type Dao struct {
	c *config.Config
}

func NewDao(conf *config.Config) (*Dao, error) {
	d := &Dao{
		c: conf,
	}
	return d, nil
}
