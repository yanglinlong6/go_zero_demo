package logic

import (
	"context"

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
	return &types.Response{"请求成功" + demoResp.Pong}, nil
}
