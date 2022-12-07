package main

import (
	"flag"
	"os"
	"os/signal"

	"github.com/hiroyaonoe/bcop-proxy/proxy/admin/api/http"
	"github.com/hiroyaonoe/bcop-proxy/proxy/admin/api/http/controller"
	"github.com/hiroyaonoe/bcop-proxy/proxy/admin/usecase"
	"github.com/hiroyaonoe/bcop-proxy/proxy/proxy"
	"github.com/hiroyaonoe/bcop-proxy/proxy/repository/inmemory"
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
	adminSrv := http.NewServer(ctrlEnv)
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
