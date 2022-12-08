package mysql

import (
	"context"

	"github.com/hiroyaonoe/bcop-proxy/controller/entity"
	"github.com/hiroyaonoe/bcop-proxy/controller/repository"
	"github.com/jmoiron/sqlx"
)

type Route struct {
	db *sqlx.DB
}

var _ repository.Route = &Route{}

func NewRoute(db *sqlx.DB) *Route {
	return &Route{db: db}
}

func (r *Route) GetWithProxyID(ctx context.Context, proxyID string) ([]entity.Route, error) {
	return nil, nil
}

func (r *Route) Upsert(ctx context.Context, routes []entity.Route) error {
	return nil
}

func (r *Route) Delete(ctx context.Context, routes []entity.Route) error {
	return nil
}
