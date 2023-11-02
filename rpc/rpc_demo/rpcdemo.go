package main

import (
	"flag"
	"fmt"

	"rpc_demo/internal/config"
	"rpc_demo/internal/server"
	"rpc_demo/internal/svc"
	"rpc_demo/rpc_demo"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/rpcdemo.yaml", "the config file")

func main() {
	// grpcui -plaintext 127.0.0.1:8080
	// goctl model mysql datasource -url="root:986203@tcp(112.74.125.238:3306)/neshield" -table="t_user_yang" -dir="./rpc/rpc_demo/sqlc/usermodel"
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		rpc_demo.RegisterRpcDemoServer(grpcServer, server.NewRpcDemoServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
