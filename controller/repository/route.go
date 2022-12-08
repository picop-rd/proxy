package repository

import (
	"context"

	"github.com/hiroyaonoe/bcop-proxy/controller/entity"
)

type Route interface {
	Get(ctx context.Context, proxyID, envID string) (entity.Route, error)
	Upsert(ctx context.Context, route entity.Route) error
	Delete(ctx context.Context, proxyID, envID string) error
}
