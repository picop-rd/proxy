package http

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/hiroyaonoe/bcop-proxy/admin/api/http/controller"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

type Router struct {
	echo *echo.Echo
	env  *controller.Env
}

func NewRouter(env *controller.Env) *Router {
	e := echo.New()

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))
	return &Router{
		echo: e,
		env:  env,
	}
}

func (r *Router) Set() {
	admin := r.echo.Group("/admin")

	admin.GET("/env/:env-id", r.env.Get)
	admin.POST("/envs", r.env.Post)
	admin.DELETE("/env/:env-id", r.env.Delete)
}

func (r *Router) Run(address string) {
	go func() {
		if err := r.echo.Start(address); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := r.echo.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("failed to shutdown")
	}
}
