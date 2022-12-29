package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/hiroyaonoe/bcop-proxy-controller/app/api/http/client"
	"github.com/hiroyaonoe/bcop-proxy/app/admin/api/http/server"
	"github.com/hiroyaonoe/bcop-proxy/app/admin/api/http/server/controller"
	"github.com/hiroyaonoe/bcop-proxy/app/admin/usecase"
	"github.com/hiroyaonoe/bcop-proxy/app/proxy"
	"github.com/hiroyaonoe/bcop-proxy/app/repository/inmemory"
	"github.com/rs/zerolog/log"
)

func main() {
	adminPort := flag.String("admin-port", "9001", "admin listen port")
	proxyPort := flag.String("proxy-port", "9000", "proxy listen port")
	propagate := flag.Bool("propagate", true, "header propagation?")
	defaultAddr := flag.String("default-addr", "", "default address")
	controllerURL := flag.String("controller-url", "http://localhost:8080", "proxy controller url")
	id := flag.String("id", "", "proxy id")

	flag.Parse()

	if len(*id) == 0 {
		log.Fatal().Msg("proxy id must exist")
	}

	repoEnv := inmemory.NewEnv()

	ucEnv := usecase.NewEnv(repoEnv)
	ctrlEnv := controller.NewEnv(ucEnv)
	adminSrv := server.NewServer(ctrlEnv)
	adminSrv.SetRoute()

	proxySrv := &proxy.Server{
		Env:                         repoEnv,
		GetEnvIDFromHeaderValueFunc: proxy.GetEnvIDFromBaggage,
		Propagate:                   *propagate,
		DefaultAddr:                 *defaultAddr,
	}

	go adminSrv.Run(":" + *adminPort)
	defer adminSrv.Close()
	go proxySrv.Start(":" + *proxyPort)
	defer proxySrv.Close()

	// Activate myself
	controllerCli := client.NewClient(http.DefaultClient, *controllerURL)
	envCli := client.NewProxy(controllerCli)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				err := envCli.Activate(ctx, *id)
				if err == nil {
					log.Info().Msg("activated")
					return
				}
				log.Error().Err(err).Msg("failed to activate myself for controller")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
