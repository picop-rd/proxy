package usecase

import (
	"context"
	"fmt"

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
	proxy.Activate = false
	if err := proxy.Validate(); err != nil {
		return fmt.Errorf("invalid proxy: %w", err)
	}

	err := p.proxy.Upsert(ctx, proxy)
	if err != nil {
		return fmt.Errorf("failed to register proxy to repository: %w", err)
	}
	return nil
}

func (p *Proxy) Activate(ctx context.Context, proxyID string) ([]entity.Route, error) {
	return nil, nil
}

func (p *Proxy) Delete(ctx context.Context, proxyID string) error {
	return nil
}
