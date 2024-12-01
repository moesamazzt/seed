package sse

import (
	"context"

	"chat/api/internal/svc"
	"chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendLogic {
	return &SendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendLogic) Send(req *types.SendReq) (resp *types.SendResp, err error) {
	// todo: add your logic here and delete this line

	return
}
