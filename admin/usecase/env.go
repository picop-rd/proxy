package usecase

import (
	"github.com/hiroyaonoe/bcop-proxy/entity"
	"github.com/hiroyaonoe/bcop-proxy/repository"
)

type Env struct {
	repo repository.Env
}

func NewEnv(repo repository.Env) *Env {
	return &Env{repo: repo}
}

func (e *Env) Get(id string) (entity.Env, error) {
	return entity.Env{}, nil
}

func (e *Env) Register(env entity.Env) error {
	return nil
}

func (e *Env) Delete(id string) error {
	return nil
}
