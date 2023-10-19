package cache

import (
	"applet-server/internal/conf"
	"github.com/redis/go-redis/v9"
)

func NewRedisCache(redisConfig *conf.Data) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:         redisConfig.Redis.Addr,
		Password:     redisConfig.Redis.Password,
		DB:           int(redisConfig.Redis.Db),
		ReadTimeout:  redisConfig.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: redisConfig.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  redisConfig.Redis.ReadTimeout.AsDuration(),
		MaxRetries:   int(redisConfig.Redis.MaxRetries),
	})
}
