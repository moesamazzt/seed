package sse

import (
	"net/http"

	"chat/api/internal/logic/sse"
	"chat/api/internal/svc"
	"chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SubscribeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SubscribeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sse.NewSubscribeLogic(r.Context(), svcCtx)
		err := l.Subscribe(w, &req, r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
