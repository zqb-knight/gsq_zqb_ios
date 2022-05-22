package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var defaultValue = map[string]string{
	"img":   "https://s1.ax1x.com/2022/05/21/OjfTPg.jpg",
	"title": "愿你如阳光，明媚不忧伤",
}

func GetRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "43.138.27.78:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

func GetKey(ctx context.Context, key string) string {
	cli := GetRedisClient()
	value, err := cli.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return defaultValue[key]
	}
	return value
}

func SetKey(ctx context.Context, key string, value string) string {
	cli := GetRedisClient()
	err := cli.Set(ctx, key, value, 0).Err()
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	return value
}
