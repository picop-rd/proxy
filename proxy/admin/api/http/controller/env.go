package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/hiroyaonoe/bcop-proxy/proxy/admin/usecase"
	"github.com/hiroyaonoe/bcop-proxy/proxy/entity"
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
		if errors.Is(err, entity.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		log.Error().Err(err).Msg("unexpected error GET /admin/env/:env-id")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, env)
}

func (e *Env) Post(c echo.Context) error {
	envs := []entity.Env{}
	if err := c.Bind(&envs); err != nil {
		log.Debug().Err(err).Msg("aaa")
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	log.Debug().Str("env", fmt.Sprintf("%#v", envs)).Send()
	err := e.uc.Register(c.Request().Context(), envs)
	if err != nil {
		if errors.Is(err, entity.ErrInvalid) {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		log.Error().Err(err).Msg("unexpected error POST /admin/envs")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusCreated)
}

func (e *Env) Delete(c echo.Context) error {
	envID := c.Param("env-id")
	if len(envID) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err := e.uc.Delete(c.Request().Context(), envID)
	if err != nil {
		log.Error().Err(err).Msg("unexpected error DELETE /admin/env/:env-id")
	}

	return c.NoContent(http.StatusOK)
}
