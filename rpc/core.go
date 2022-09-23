package main

import (
	"flag"
	"fmt"

	"github.com/suyuan32/simple-admin-core/rpc/internal/config"
	"github.com/suyuan32/simple-admin-core/rpc/internal/server"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-tools/plugins/registry/consul"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/core.yaml", "the config file")

func main() {
	flag.Parse()

	var consulConfig config.ConsulConfig
	conf.MustLoad(*configFile, &consulConfig)

	var c config.Config
	client, err := consulConfig.Consul.NewClient()
	logx.Must(err)
	consul.LoadYAMLConf(client, "coreRpcConf", &c)

	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		core.RegisterCoreServer(grpcServer, server.NewCoreServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	err = consul.RegisterService(consulConfig.Consul)
	logx.Must(err)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
