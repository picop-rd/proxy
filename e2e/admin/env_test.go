package admin

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/hiroyaonoe/bcop-proxy/app/admin/api/http/client"
	"github.com/hiroyaonoe/bcop-proxy/app/admin/api/http/server"
	"github.com/hiroyaonoe/bcop-proxy/app/admin/api/http/server/controller"
	"github.com/hiroyaonoe/bcop-proxy/app/admin/usecase"
	"github.com/hiroyaonoe/bcop-proxy/app/entity"
	"github.com/hiroyaonoe/bcop-proxy/app/repository/inmemory"
)

func prepareServer(t *testing.T) (string, func(), error) {
	t.Helper()

	repoEnv := inmemory.NewEnv()

	ucEnv := usecase.NewEnv(repoEnv)
	ctrlEnv := controller.NewEnv(ucEnv)

	srv := server.NewServer(ctrlEnv)
	srv.SetRoute()

	go srv.Run(":0")

	for i := 0; i < 50; i++ {
		if srv.Echo.Listener == nil {
			time.Sleep(100 * time.Millisecond)
			continue
		}
		return srv.Echo.ListenerAddr().String(), srv.Close, nil
	}
	return "", nil, fmt.Errorf("failed to get listen addr")
}

func TestEnv_Senario_Normal1(t *testing.T) {
	t.Parallel()
	addr, closer, err := prepareServer(t)
	if err != nil {
		t.Fatalf("failed to listen server: %s", err)
	}
	defer closer()

	cli := client.NewClient(&http.Client{}, fmt.Sprintf("http://%s", addr))
	envCli := client.NewEnv(cli)

	envA := entity.Env{
		EnvID:       "A",
		Destination: "destination:portA",
	}
	envB := entity.Env{
		EnvID:       "B",
		Destination: "destination:portB",
	}

	t.Run("新しいEnvを登録できる", func(t *testing.T) {
		err = envCli.Register(context.Background(), []entity.Env{envA, envB})
		if err != nil {
			t.Errorf("Env.Register(): error = %v", err)
		}
	})

	t.Run("登録済みのEnvを取得できる", func(t *testing.T) {
		gotEnvA, err := envCli.Get(context.Background(), "A")
		if err != nil {
			t.Errorf("Env.Get(): error = %v", err)
		}
		if diff := cmp.Diff(envA, gotEnvA); diff != "" {
			t.Errorf("Env.Get(): mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("登録済みのEnvのDestinationを変更できる", func(t *testing.T) {
		modifiedEnvA := entity.Env{
			EnvID:       "A",
			Destination: "modifiedDestination:portA",
		}
		err = envCli.Register(context.Background(), []entity.Env{modifiedEnvA})
		if err != nil {
			t.Errorf("Env.Register(): error = %v", err)
		}
		gotModifiedEnvA, err := envCli.Get(context.Background(), "A")
		if err != nil {
			t.Errorf("Env.Get(): error = %v", err)
		}
		if diff := cmp.Diff(modifiedEnvA, gotModifiedEnvA); diff != "" {
			t.Errorf("Env.Get(): mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("Delete: 登録済みのEnvを削除できる", func(t *testing.T) {
		err = envCli.Delete(context.Background(), "A")
		if err != nil {
			t.Errorf("Env.Delete(): delete: error = %v", err)
		}
		gotDeletedEnvA, err := envCli.Get(context.Background(), "A")
		if err == nil {
			t.Errorf("Env.Get(): delete: non-nil error: gotEnv: %v", gotDeletedEnvA)
		}
	})
}

func TestEnv_Senario_TooManyRequest(t *testing.T) {
	t.Parallel()
	addr, closer, err := prepareServer(t)
	if err != nil {
		t.Fatalf("failed to listen server: %s", err)
	}
	defer closer()

	cli := client.NewClient(&http.Client{}, fmt.Sprintf("http://%s", addr))
	envCli := client.NewEnv(cli)

	envA := entity.Env{
		EnvID:       "A",
		Destination: "destination:portA",
	}
	envB := entity.Env{
		EnvID:       "B",
		Destination: "destination:portB",
	}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := envCli.Register(context.Background(), []entity.Env{envA, envB})
			if err != nil {
				t.Errorf("Env.Register(): error = %v", err)
			}
			_, err = envCli.Get(context.Background(), "B")
			if err != nil {
				t.Errorf("Env.Get(): error = %v", err)
			}
			err = envCli.Delete(context.Background(), "A")
			if err != nil {
				t.Errorf("Env.Delete(): error = %v", err)
			}
		}()
	}
	wg.Wait()
}
