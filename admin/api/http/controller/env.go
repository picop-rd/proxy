package controller

import (
	"net/http"

	"github.com/hiroyaonoe/bcop-proxy/admin/usecase"
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
	envID := c.Param("env-id")
	if len(envID) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	env, err := e.uc.Get(c.Request().Context(), envID)
	if err != nil {
		log.Error().Err(err).Msg("unexpected error GET /admin/env/:env-id")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, env)
}
