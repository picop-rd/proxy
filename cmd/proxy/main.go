package main

import (
	"flag"
	"os"
	"os/signal"

	"github.com/hiroyaonoe/bcop-proxy/app/admin/api/http/server"
	"github.com/hiroyaonoe/bcop-proxy/app/admin/api/http/server/controller"
	"github.com/hiroyaonoe/bcop-proxy/app/admin/usecase"
	"github.com/hiroyaonoe/bcop-proxy/app/proxy"
	"github.com/hiroyaonoe/bcop-proxy/app/repository/inmemory"
)

func main() {
	adminPort := flag.String("admin-port", "9001", "admin listen port")

	proxyPort := flag.String("proxy-port", "9000", "proxy listen port")
	propagate := flag.Bool("propagate", true, "header propagation?")
	defaultAddr := flag.String("default-addr", "", "default address")

	flag.Parse()

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

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
