package main

import (
	"flag"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"

	"github.com/suyuan32/simple-admin-core/job/internal/config"
	"github.com/suyuan32/simple-admin-core/job/internal/listen"
)

var configFile = flag.String("f", "etc/core.yaml", "Specify the config file")

func main() {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFile, &c)

	if err := c.SetUp(); err != nil {
		panic(err)
	}

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	for _, mq := range listen.Mqs(c) {
		serviceGroup.Add(mq)
	}

	serviceGroup.Start()

	// keep running
	select {}

}
