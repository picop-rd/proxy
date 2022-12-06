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
	envs, err := e.uc.Get(c.Request().Context())
	if err != nil {
		log.Error().Err(err).Msg("unexpected error GET /admin/env")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, envs)
}
