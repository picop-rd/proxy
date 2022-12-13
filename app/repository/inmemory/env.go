package inmemory

import (
	"context"

	"github.com/hiroyaonoe/bcop-proxy/app/entity"
	"github.com/hiroyaonoe/bcop-proxy/app/repository"
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
	env, ok := e.db[id]
	if !ok {
		return entity.Env{}, entity.ErrNotFound
	}
	if err := env.Validate(); err != nil {
		return entity.Env{}, entity.ErrNotFound
	}
	return env, nil
}

func (e *Env) Upsert(_ context.Context, envs []entity.Env) error {
	for _, v := range envs {
		e.db[v.EnvID] = v
	}
	return nil
}

func (e *Env) Delete(_ context.Context, id string) error {
	e.db[id] = entity.Env{}
	return nil
}
