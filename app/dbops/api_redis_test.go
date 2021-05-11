package dbops

import (
	"fmt"
	"testing"
)

func TestRedisConn(t *testing.T) {
	t.Run("redis", testRedis)
}

func testRedis(t *testing.T) {
	err := redisConn.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := redisConn.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("key: %v\n", val)
}