package usecase

import (
	"context"

	"github.com/hiroyaonoe/bcop-proxy/controller/entity"
	"github.com/hiroyaonoe/bcop-proxy/controller/repository"
)

type Proxy struct {
	proxy repository.Proxy
}

func NewProxy(proxy repository.Proxy) *Proxy {
	return &Proxy{proxy: proxy}
}

func (p *Proxy) Register(ctx context.Context, proxy entity.Proxy) error {
	return nil
}

func (p *Proxy) Activate(ctx context.Context, proxyID string) ([]entity.Route, error) {
	return nil, nil
}

func (p *Proxy) Delete(ctx context.Context, proxyID string) error {
	return nil
}
