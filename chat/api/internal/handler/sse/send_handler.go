package sse

import (
	"net/http"

	"chat/api/internal/logic/sse"
	"chat/api/internal/svc"
	"chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sse.NewSendLogic(r.Context(), svcCtx)
		resp, err := l.Send(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
