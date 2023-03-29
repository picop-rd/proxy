package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/picop-rd/proxy/app/entity"
)

type Env struct {
	client *Client
}

func NewEnv(client *Client) *Env {
	return &Env{
		client: client,
	}
}

func (e *Env) Get(ctx context.Context, envID string) (entity.Env, error) {
	resp, err := e.client.Get(ctx, []string{"admin", "env", envID})
	if err != nil {
		return entity.Env{}, fmt.Errorf("proxy admin client: Get: failed to request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return entity.Env{}, fmt.Errorf("proxy admin client: Get: response status code is not 200, status: %s", resp.Status)
	}
	dec := json.NewDecoder(resp.Body)
	var env entity.Env
	err = dec.Decode(&env)
	if err != nil {
		return entity.Env{}, fmt.Errorf("proxy admin client: Get: failed to decode json from response body: %w", err)
	}
	return env, nil
}

func (e *Env) Register(ctx context.Context, envs []entity.Env) error {
	byteBody, err := json.Marshal(envs)
	if err != nil {
		return fmt.Errorf("proxy admin client: Register: failed to encode json from envs: %w", err)
	}
	resp, err := e.client.Put(ctx, []string{"admin", "envs"}, "application/json", bytes.NewReader(byteBody))
	if err != nil {
		return fmt.Errorf("proxy admin client: Register: failed to request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("proxy admin client: Register: response status code is not 200, status: %s", resp.Status)
	}
	return nil
}

func (e *Env) Delete(ctx context.Context, envID string) error {
	resp, err := e.client.Delete(ctx, []string{"admin", "env", envID})
	if err != nil {
		return fmt.Errorf("proxy admin client: Delete: failed to request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("proxy admin client: Delete: response status code is not 200, status: %s", resp.Status)
	}
	return nil
}
