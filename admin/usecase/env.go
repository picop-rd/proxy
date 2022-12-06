package usecase

import (
	"context"

	"github.com/hiroyaonoe/bcop-proxy/entity"
	"github.com/hiroyaonoe/bcop-proxy/repository"
)

type Env struct {
	repo repository.Env
}

func NewEnv(repo repository.Env) *Env {
	return &Env{repo: repo}
}

func (e *Env) Get(ctx context.Context, id string) (entity.Env, error) {
	return entity.Env{}, nil
}

func (e *Env) Register(ctx context.Context, env entity.Env) error {
	return nil
}

func (e *Env) Delete(ctx context.Context, id string) error {
	return nil
}
