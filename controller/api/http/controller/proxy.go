package controller

import (
	"github.com/hiroyaonoe/bcop-proxy/controller/usecase"
	echo "github.com/labstack/echo/v4"
)

type Proxy struct {
	uc *usecase.Proxy
}

func NewProxy(uc *usecase.Proxy) *Proxy {
	return &Proxy{uc: uc}
}

func (p *Proxy) Register(c echo.Context) error {
	return nil
}

func (p *Proxy) Activate(c echo.Context) error {
	return nil
}

func (p *Proxy) Delete(c echo.Context) error {
	return nil
}
