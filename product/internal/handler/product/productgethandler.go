package product

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"product/internal/logic/product"
	"product/internal/svc"
	"product/internal/types"
)

func ProductGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductGetReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := product.NewProductGetLogic(r.Context(), svcCtx)
		resp, err := l.ProductGet(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
