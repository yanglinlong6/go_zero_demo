package svc

import (
	"api_demo/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"rpc_demo/rpcdemoclient"
	"rpc_demo/sqlc/usermodel"
)

type ServiceContext struct {
	Config     config.Config
	DemoRpc    rpcdemoclient.RpcDemo
	UsersModel usermodel.TUserYangModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.MySQL.DataSource)
	return &ServiceContext{
		Config:     c,
		DemoRpc:    rpcdemoclient.NewRpcDemo(zrpc.MustNewClient(c.DemoRpcConf)),
		UsersModel: usermodel.NewTUserYangModel(mysqlConn),
	}
}
