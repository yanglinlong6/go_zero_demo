package logic

import (
	"context"
	"database/sql"
	"rpc_demo/sqlc/usermodel"

	"api_demo/internal/svc"
	"api_demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	pb "rpc_demo/rpc_demo"
)

type Api_demoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApi_demoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Api_demoLogic {
	return &Api_demoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Api_demoLogic) Api_demo(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	demoResp, err := l.svcCtx.DemoRpc.Ping(l.ctx, &pb.Request{Ping: "yang ping"})
	l.Logger.Info(demoResp)
	yangParam := &usermodel.TUserYang{}
	yangParam.Username = sql.NullString{String: "text api", Valid: true}
	yangParam.Password = sql.NullString{String: "dfsadsafdsf", Valid: true}
	insert, _ := l.svcCtx.UsersModel.Insert(l.ctx, yangParam)
	affected, _ := insert.RowsAffected()
	l.Logger.Info(affected)
	return &types.Response{"请求成功" + demoResp.Pong}, nil
}
