package repository

import (
	"context"

	"github.com/hiroyaonoe/bcop-proxy/entity"
)

type Env interface {
	Get(ctx context.Context) ([]entity.Env, error)
	Upsert(ctx context.Context, env []entity.Env) error
	Delete(ctx context.Context, ids []string) error
}
