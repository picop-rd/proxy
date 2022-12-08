package mysql

import (
	"context"

	"github.com/hiroyaonoe/bcop-proxy/controller/entity"
	"github.com/hiroyaonoe/bcop-proxy/controller/repository"
)

type Proxy struct {
	db *DB
}

var _ repository.Proxy = &Proxy{}

func NewProxy(db *DB) *Proxy {
	return &Proxy{db: db}
}

func (p *Proxy) Get(ctx context.Context, proxyID string) (entity.Proxy, error) {
	return entity.Proxy{}, nil
}

func (p *Proxy) Upsert(ctx context.Context, proxy entity.Proxy) error {
	return nil
}

func (p *Proxy) Delete(ctx context.Context, proxyID string) error {
	return nil
}
