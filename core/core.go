package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/sjxiang/go-zero-cloud-disk/core/internal/config"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/handler"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"

	"github.com/joho/godotenv"
)

var configFile = flag.String("f", "etc/core-api.yaml", "the config file")

func Init() {
	// 从本地读取环境变量
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	flag.Parse()

	Init()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
