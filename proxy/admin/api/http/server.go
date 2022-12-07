package http

import (
	"context"
	"net/http"
	"time"

	"github.com/hiroyaonoe/bcop-proxy/proxy/admin/api/http/controller"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

type Server struct {
	echo *echo.Echo
	env  *controller.Env
}

func NewServer(env *controller.Env) *Server {
	e := echo.New()

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info().
				Str("Method", v.Method).
				Str("URI", v.URI).
				Int("Status", v.Status).
				Msg("Request")

			return nil
		},
	}))
	return &Server{
		echo: e,
		env:  env,
	}
}

func (s *Server) SetRoute() {
	admin := s.echo.Group("/admin")

	admin.GET("/env/:env-id", s.env.Get)
	admin.PUT("/envs", s.env.Post)
	admin.DELETE("/env/:env-id", s.env.Delete)
}

func (s *Server) Run(address string) {
	if err := s.echo.Start(address); err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("shutting down the server")
	}
}

func (s *Server) Close() {
	log.Info().Msg("admin shutdown")
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.echo.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("failed to shutdown")
	}
}
