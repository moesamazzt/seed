package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"user/api/internal/config"
	"user/api/internal/middleware"
)

type ServiceContext struct {
	Config          config.Config
	AuthInterceptor rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		AuthInterceptor: middleware.NewAuthInterceptorMiddleware().Handle,
	}
}
