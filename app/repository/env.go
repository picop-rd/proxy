package repository

import (
	"context"

	"github.com/picop-rd/proxy/app/entity"
)

type Env interface {
	Get(ctx context.Context, id string) (entity.Env, error)
	Upsert(ctx context.Context, env []entity.Env) error
	Delete(ctx context.Context, id string) error
}
