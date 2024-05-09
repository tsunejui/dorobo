package api

import (
	"dorobo/configs"
	"dorobo/lib/process"
)

type Controller struct {
	version string
	pm      *process.Manager
}

const (
	defaultVersion = "v0.0.1"
)

func New(conf configs.WebConfig) (*Controller, error) {
	v := defaultVersion
	if conf.Version != "" {
		v = conf.Version
	}
	return &Controller{
		version: v,
	}, nil
}

func (c *Controller) Close() error {
	return nil
}

func (c *Controller) WithPM(pm *process.Manager) *Controller {
	c.pm = pm
	return c
}
