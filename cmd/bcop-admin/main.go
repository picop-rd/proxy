package main

import (
	"flag"

	"github.com/hiroyaonoe/bcop-proxy/admin/api/http"
	"github.com/hiroyaonoe/bcop-proxy/admin/api/http/controller"
	"github.com/hiroyaonoe/bcop-proxy/admin/usecase"
	"github.com/hiroyaonoe/bcop-proxy/repository/inmemory"
)

func main() {
	port := flag.String("port", "8080", "listen port")
	flag.Parse()

	repoEnv := inmemory.NewEnv()
	ucEnv := usecase.NewEnv(repoEnv)
	ctrlEnv := controller.NewEnv(ucEnv)
	router := http.NewRouter(ctrlEnv)
	router.Set()

	router.Run(*port)
}
