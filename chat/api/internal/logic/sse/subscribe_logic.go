package sse

import (
	"chat/api/internal/utils"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"chat/api/internal/svc"
	"chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubscribeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubscribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubscribeLogic {
	return &SubscribeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubscribeLogic) Subscribe(w http.ResponseWriter, req *types.SubscribeReq, r *http.Request) error {
	// 校验用户 ID
	if req.UserId == "" {
		return errors.New("user ID cannot be empty")
	}

	// 将用户加入订阅队列（可以是 Redis、数据库等）
	err := utils.AddToSubscriptionQueue(req.UserId)
	if err != nil {
		return err
	}

	// 设置 HTTP 头，表明这是一个 SSE 连接
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// 将连接转化为 http.Flusher 类型
	flusher, ok := w.(http.Flusher)
	if !ok {
		return errors.New("unable to convert response writer to http.Flusher")
	}

	// 进入一个循环，持续推送消息给订阅的用户
	for {
		// 从 Redis 等地方获取消息，这里以模拟数据为例
		message := utils.GetLatestMessageForUser(req.UserId)
		if message != "" {
			// 推送消息到客户端
			_, err := fmt.Fprintf(w, "data: %s\n\n", message)
			if err != nil {
				logx.Errorf("failed to send message: %v", err)
				return err
			}
			flusher.Flush() // 强制刷新输出流，确保客户端接收到消息
		}

		// 每秒检查客户端连接状态
		select {
		case <-time.After(1 * time.Second): // 每秒发送一次数据
			// 检查客户端连接是否已关闭
			if r.Context().Err() != nil {
				// 如果客户端断开连接，退出循环
				logx.Info("client disconnected")
				return nil
			}
		}
	}
}
