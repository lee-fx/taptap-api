package dbops

import (
	"api/app/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var (
	redisConn *redis.Client
	ctx = context.Background()
)

func init() {

	addr := config.GetRedisAddr() + ":" + config.GetRedisPort()

	password := config.GetRedisPassword()
	db := config.GetRedisAppDb()
	redisConn = redis.NewClient(&redis.Options{
		Addr:     addr, // 要连接的redis IP:port
		Password: password,               // redis 密码
		DB:       db,                // 要连接的redis 库
	})
	// 检测心跳
	_, err := redisConn.Ping(ctx).Result()
	if err != nil {
		fmt.Println("connect redis failed")
	}
	//fmt.Printf("redis ping result: %s\n", pong)
}
