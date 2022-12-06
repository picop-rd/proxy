package usecase

import (
	"context"
	"fmt"

	"github.com/hiroyaonoe/bcop-proxy/entity"
	"github.com/hiroyaonoe/bcop-proxy/repository"
)

type Env struct {
	repo repository.Env
}

func NewEnv(repo repository.Env) *Env {
	return &Env{repo: repo}
}

func (e *Env) Get(ctx context.Context) ([]entity.Env, error) {
	envs, err := e.repo.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get envs from repository: %w", err)
	}
	return envs, nil
}

func (e *Env) Register(ctx context.Context, env entity.Env) error {
	return nil
}

func (e *Env) Delete(ctx context.Context, id string) error {
	return nil
}
