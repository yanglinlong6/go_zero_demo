package svc

import (
	"api_demo/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
	"rpc_demo/rpcdemoclient"
)

type ServiceContext struct {
	Config  config.Config
	DemoRpc rpcdemoclient.RpcDemo
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		DemoRpc: rpcdemoclient.NewRpcDemo(zrpc.MustNewClient(c.DemoRpcConf)),
	}
}
