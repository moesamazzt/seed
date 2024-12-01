// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"
	"time"

	account "user/api/internal/handler/account"
	profile "user/api/internal/handler/profile"
	"user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/login",
				Handler: account.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/refresh",
				Handler: account.RefreshHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: account.RegisterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/user/account"),
		rest.WithTimeout(10000*time.Millisecond),
		rest.WithMaxBytes(1048576),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthInterceptor},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/base",
					Handler: profile.BaseHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/user/profile"),
		rest.WithTimeout(10000*time.Millisecond),
		rest.WithMaxBytes(1048576),
	)
}