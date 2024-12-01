package utils

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var Rdb *redis.Client

func InitRedis(addr string, password string, db int) error {
	// 创建 Redis 客户端
	Rdb = redis.NewClient(&redis.Options{
		Addr:     addr,     // Redis 地址，例如 "localhost:6379"
		Password: password, // Redis 密码，如果没有密码留空
		DB:       db,       // 使用的 Redis 数据库索引，默认为 0
	})

	// 测试连接
	ctx := context.Background()
	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		log.Printf("Redis connection failed: %v", err)
		return err
	}

	log.Println("Redis connected successfully")
	return nil
}

// AddToSubscriptionQueue 将用户加入订阅队列
func AddToSubscriptionQueue(userId string) error {
	if userId == "" {
		return fmt.Errorf("userId cannot be empty")
	}
	// 将用户加入某个订阅队列
	return Rdb.SAdd(context.Background(), "subscribers", userId).Err()
}

// GetLatestMessageForUser 获取某个用户的最新消息（可以从 Redis 中获取）
func GetLatestMessageForUser(userId string) string {
	// 这里简单模拟获取消息，你可以从 Redis 或数据库中获取实际的消息
	// 假设每次返回一个时间戳作为模拟消息
	return fmt.Sprintf("New message for %s at %s", userId, time.Now().Format(time.RFC3339))
}
