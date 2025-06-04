package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

const URL = "redis://:123456@127.0.0.1:6379/0"

func main() {
	opt, err := redis.ParseURL(URL)
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(opt)

	ctx := context.Background()
	rdb.Set(ctx, "hello", "world", redis.KeepTTL)

	value := rdb.Get(ctx, "hello").Val()
	fmt.Println(value)
}
