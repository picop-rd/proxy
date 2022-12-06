package controller

import (
	"net/http"

	"github.com/hiroyaonoe/bcop-proxy/admin/usecase"
	"github.com/hiroyaonoe/bcop-proxy/entity"
	echo "github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type Env struct {
	uc *usecase.Env
}

func NewEnv(uc *usecase.Env) *Env {
	return &Env{uc: uc}
}

func (e *Env) Get(c echo.Context) error {
	envs, err := e.uc.Get(c.Request().Context())
	if err != nil {
		log.Error().Err(err).Msg("unexpected error GET /admin/env")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, envs)
}

func (e *Env) Post(c echo.Context) error {
	envs := []entity.Env{}
	if err := c.Bind(envs); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err := e.uc.Register(c.Request().Context(), envs)
	if err != nil {
		log.Error().Err(err).Msg("unexpected error POST /admin/envs")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusCreated)
}
