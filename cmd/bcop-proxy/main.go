package main

import (
	"flag"

	"github.com/hiroyaonoe/bcop-proxy/proxy"
	"github.com/hiroyaonoe/bcop-proxy/repository/inmemory"
)

func main() {
	propagate := flag.Bool("propagate", true, "header propagation?")
	defalutAddr := flag.String("default", "", "defalut address")
	port := flag.String("port", "8081", "listen port")
	flag.Parse()

	envRepo := inmemory.NewEnv()
	server := &proxy.Server{
		Env:                         envRepo,
		GetEnvIDFromHeaderValueFunc: proxy.GetEnvIDFromBaggage,
		Propagate:                   *propagate,
		DefaultAddr:                 *defalutAddr,
	}

	server.Start(":" + *port)

}
