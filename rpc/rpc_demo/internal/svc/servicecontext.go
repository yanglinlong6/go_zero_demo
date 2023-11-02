package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"rpc_demo/internal/config"
)
import "rpc_demo/sqlc/usermodel"

type ServiceContext struct {
	Config     config.Config
	UsersModel usermodel.TUserYangModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.MySQL.DataSource)
	return &ServiceContext{
		Config:     c,
		UsersModel: usermodel.NewTUserYangModel(mysqlConn),
	}
}
