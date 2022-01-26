package lib

import (
	"github.com/david8998/lib/config"
	"github.com/go-redis/redis/v8"
	"time"
)

func GetRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.C.Redis.Addr,
		Password: config.C.Redis.Password,
		DB:       config.C.Redis.DB,

		//超时配置
		DialTimeout:  config.C.Redis.DialTimeout * time.Millisecond,
		ReadTimeout:  config.C.Redis.ReadTimeout * time.Millisecond,
		WriteTimeout: config.C.Redis.WriteTimeout * time.Millisecond,
		MaxConnAge:   config.C.Redis.MaxConnAge * time.Second,

		//连接数配置
		PoolSize:     config.C.Redis.MaxCons,
		MinIdleConns: config.C.Redis.MinIdleCons,
	})
	return client
}
