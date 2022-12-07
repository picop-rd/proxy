package main

import (
	"flag"
	"sync"

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
	defaultAddr := flag.String("default-addr", "", "defalut address")

	flag.Parse()

	repoEnv := inmemory.NewEnv()

	ucEnv := usecase.NewEnv(repoEnv)
	ctrlEnv := controller.NewEnv(ucEnv)
	router := http.NewRouter(ctrlEnv)
	router.Set()

	server := &proxy.Server{
		Env:                         repoEnv,
		GetEnvIDFromHeaderValueFunc: proxy.GetEnvIDFromBaggage,
		Propagate:                   *propagate,
		DefaultAddr:                 *defaultAddr,
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go router.Run(":" + *adminPort)
	go server.Start(":" + *proxyPort)
	wg.Wait()
}
