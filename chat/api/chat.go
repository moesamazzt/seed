package main

import (
	"chat/api/internal/utils"
	"flag"
	"fmt"
	"log"

	"chat/api/internal/config"
	"chat/api/internal/handler"
	"chat/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/chat.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 初始化 Redis，指定地址、密码和数据库
	err := utils.InitRedis("localhost:6379", "v2space", 0)
	if err != nil {
		log.Fatalf("Failed to initialize Redis: %v", err)
	}

	server := rest.MustNewServer(c.RestConf, rest.WithCors("*"))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
