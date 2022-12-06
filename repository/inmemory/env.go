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
	return entity.Env{}, nil
}

func (e *Env) Upsert(_ context.Context, env entity.Env) error {
	return nil
}

func (e *Env) Delete(_ context.Context, id string) error {
	return nil
}
