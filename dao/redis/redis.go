package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"web_app/settings"
)

var rdb *redis.Client

func Init(conf *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.Host, strconv.Itoa(conf.Port)),
		Password: conf.Password,
		DB:       conf.DB,
		PoolSize: conf.PoolSize,
	})

	_, err = rdb.Ping().Result()

	return err
}

func Close() {
	rdb.Close()
}
