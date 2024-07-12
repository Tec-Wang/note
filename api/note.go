package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/service"

	"wangzheng/brain/api/internal/config"
	"wangzheng/brain/api/internal/handler"
	"wangzheng/brain/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"looklook/app/order/cmd/mq/internal/listen"
)

var configFile = flag.String("f", "etc/note-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	for _, mq := range listen.Mqs(c) {
		serviceGroup.Add(mq)
	}

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
