package controller

import (
	"github.com/hiroyaonoe/bcop-proxy/admin/usecase"
	echo "github.com/labstack/echo/v4"
)

type Env struct {
	uc *usecase.Env
}

func NewEnv(uc *usecase.Env) *Env {
	return &Env{uc: uc}
}

func (e *Env) Get(ctx echo.Context) error {
	return nil
}
