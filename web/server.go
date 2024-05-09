package web

import (
	"context"
	"dorobo/api"
	"dorobo/configs"
	"dorobo/lib/process"
	"dorobo/pkg/keep"
	"dorobo/pkg/update"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

var version string = "v0.0.1"

const (
	defaultListenPort = "8090"
)

type Server struct {
	svr *echo.Echo
	pm  *process.Manager
}

func New() (*Server, error) {
	s := &Server{
		svr: echo.New(),
	}
	if err := s.init(); err != nil {
		return nil, fmt.Errorf("failed to init server: %v", err)
	}
	return s, nil
}

func (s *Server) init() error {
	fmt.Printf("\nweb server version: %s\n", version)
	// API controller
	e := s.svr
	e.HideBanner = true
	controller, err := api.New(configs.WebConfig{
		ListenPort: defaultListenPort,
		Version:    version,
	})
	if err != nil {
		return fmt.Errorf("failed to get controller: %v", err)
	}

	// process keeper
	keeper := keep.New(controller, e)
	keeper.SetKeep(true)

	// process updater
	updater := update.New()

	// process manager
	pm := process.New(keeper, updater)
	s.pm = pm
	controller.WithPM(pm)

	// set routes
	e.GET("/hello", controller.SayHello)
	e.GET("/version", controller.GetVersion)
	e.POST("/upgrade", controller.Upgrade)
	return nil
}

func (s *Server) Run() {
	e := s.svr
	go func() {
		if err := e.Start(
			fmt.Sprintf(":%s", defaultListenPort),
		); err != http.ErrServerClosed {
			e.Logger.Fatal(err)
		}
	}()
	s.pm.RunKeeper()
}

func (s *Server) Close() error {
	ctx := context.Background()
	if err := s.svr.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to close echo server: %v", err)
	}
	return nil
}
