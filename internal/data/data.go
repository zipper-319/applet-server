package data

import (
	"applet-server/internal/data/cache"
	"applet-server/internal/data/s3"
	"github.com/redis/go-redis/v9"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, s3.NewS3Service, cache.NewRedisCache)

// Data .
type Data struct {
	S3          *s3.S3Service
	RedisClient *redis.Client
}

// NewData .
func NewData(s3 *s3.S3Service, rdb *redis.Client) (*Data, error) {
	return &Data{
		S3:          s3,
		RedisClient: rdb,
	}, nil
}
