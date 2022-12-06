package inmemory

import (
	"context"

	"github.com/hiroyaonoe/bcop-proxy/entity"
	"github.com/hiroyaonoe/bcop-proxy/repository"
)

type Env struct {
	db map[string]entity.Env
}

var _ repository.Env = &Env{}

func NewEnv() *Env {
	return &Env{
		db: map[string]entity.Env{},
	}
}

func (e *Env) Get(_ context.Context, id string) (entity.Env, error) {
	return e.db[id], nil
}

func (e *Env) Upsert(_ context.Context, envs []entity.Env) error {
	for _, v := range envs {
		e.db[v.EnvID] = v
	}
	return nil
}

func (e *Env) Delete(_ context.Context, ids []string) error {
	return nil
}
