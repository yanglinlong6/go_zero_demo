package logic

import (
	"context"
	"database/sql"
	"rpc_demo/sqlc/usermodel"

	"rpc_demo/internal/svc"
	"rpc_demo/rpc_demo"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *rpc_demo.Request) (*rpc_demo.Response, error) {
	// todo: add your logic here and delete this line
	yangParam := &usermodel.TUserYang{}
	yangParam.Username = sql.NullString{String: "text", Valid: true}
	yangParam.Password = sql.NullString{String: "dfsadsafdsf", Valid: true}
	yangParam.Remark = sql.NullString{String: "Remark", Valid: true}
	insert, _ := l.svcCtx.UsersModel.Insert(l.ctx, yangParam)
	affected, _ := insert.RowsAffected()
	l.Logger.Info(affected)
	return &rpc_demo.Response{Pong: "Pong" + in.Ping + ";"}, nil
}
