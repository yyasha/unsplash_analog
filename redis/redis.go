package redis

import (
	"context"
	"unsplash_analog/config"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisCache struct {
	rdb *redis.Client
}

var RDB RedisCache

func InitRedis() {
	RDB.rdb = redis.NewClient(&redis.Options{
		Addr:     config.Conf.RDB_ADDR,
		Password: config.Conf.RDB_PASSWORD, // no password set
		DB:       0,                        // use default DB
	})
}
