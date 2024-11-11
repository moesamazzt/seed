package product

import (
	"context"

	"product/internal/svc"
	"product/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductGetLogic {
	return &ProductGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductGetLogic) ProductGet(req *types.ProductGetReq) (resp *types.ResultProductGetResp, err error) {
	// todo: add your logic here and delete this line

	return
}
