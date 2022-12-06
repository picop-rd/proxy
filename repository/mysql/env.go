package mysql

import (
	"context"
	"fmt"

	"github.com/hiroyaonoe/bcop-proxy/entity"
	"github.com/hiroyaonoe/bcop-proxy/repository"
)

type Env struct {
	db *DB
}

var _ repository.Env = &Env{}

func NewEnv(db *DB) *Env {
	return &Env{db: db}
}

func (e *Env) Get(ctx context.Context, id string) (entity.Env, error) {
	var env entity.Env
	err := e.db.GetContext(ctx, env, "SELECT (env-id, destination) FROM env-id = $1", id)
	if err != nil {
		return entity.Env{}, fmt.Errorf("failed to get env from mysql: %w", err)
	}
	return env, nil
}

func (e *Env) Upsert(ctx context.Context, env entity.Env) error {
	return nil
}

func (e *Env) Delete(ctx context.Context, id string) error {
	return nil
}
