//		Simple Admin
//
//		This is simple admin api doc
//
//		Schemes: http, https
//		Host: localhost:8500
//		BasePath: /
//		Version: 0.0.1
//		Contact: yuansu.china.work@gmail.com
//		securityDefinitions:
//		  Token:
//		    type: apiKey
//		    name: Authorization
//		    in: header
//		security:
//		  - Token: []
//	    Consumes:
//		  - application/json
//
//		Produces:
//		  - application/json
//
// swagger:meta
package main

import (
	"flag"
	"fmt"
	"github.com/suyuan32/simple-admin-tools/plugins/registry/consul"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/api/internal/config"
	"github.com/suyuan32/simple-admin-core/api/internal/handler"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	_ "github.com/suyuan32/simple-admin-tools/plugins/registry/consul"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/core.yaml", "the config file")

func main() {
	flag.Parse()

	var consulConfig config.ConsulConfig
	conf.MustLoad(*configFile, &consulConfig)

	var c config.Config
	client, err := consulConfig.Consul.NewClient()
	logx.Must(err)
	consul.LoadYAMLConf(client, "coreApiConf", &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors("*"))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	err = consul.RegisterService(consulConfig.Consul)
	logx.Must(err)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
