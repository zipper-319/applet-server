package cache

import (
	"CBEC_backend/internal/conf"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"os"
	"testing"
)

var bc conf.Bootstrap

func init() {
	fmt.Println(os.Getwd())
	c := config.New(
		config.WithSource(
			file.NewSource("../../../configs"),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
}

func TestSetKey(t *testing.T) {
	redisClient := NewRedisCache(bc.Data)
	result := redisClient.Set(context.Background(), "hello", "", -1)
	fmt.Println(result)
}

func TestGetKey(t *testing.T) {
	redisClient := NewRedisCache(bc.Data)
	result := redisClient.Get(context.Background(), "hello")
	fmt.Println(result.Result())
}
