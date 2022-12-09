package controller

import (
	"errors"
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
	envID := c.Param("env-id")
	if len(envID) == 0 {
		log.Debug().Msg("illegal param")
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

func (e *Env) Put(c echo.Context) error {
	envs := []entity.Env{}
	if err := c.Bind(&envs); err != nil {
		log.Debug().Err(err).Msg("illegal body")
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err := e.uc.Register(c.Request().Context(), envs)
	if err != nil {
		if errors.Is(err, entity.ErrInvalid) {
			log.Debug().Err(err).Msg("illegal envs")
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		log.Error().Err(err).Msg("unexpected error POST /admin/envs")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

func (e *Env) Delete(c echo.Context) error {
	envID := c.Param("env-id")
	if len(envID) == 0 {
		log.Debug().Msg("illegal param")
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err := e.uc.Delete(c.Request().Context(), envID)
	if err != nil {
		log.Error().Err(err).Msg("unexpected error DELETE /admin/env/:env-id")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
