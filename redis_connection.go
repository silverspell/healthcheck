package healthcheckmodule

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisConnection struct {
	Host string
	User string
	Pass string
	Db   int
}

func (r *RedisConnection) Connect() error {
	r.Db, _ = strconv.Atoi(os.Getenv("REDIS_DB"))
	r.Host = os.Getenv("REDIS_HOST")
	r.User = os.Getenv("REDIS_USER")
	r.Pass = os.Getenv("REDIS_PASS")

	rdb := redis.NewClient(&redis.Options{
		Addr:        r.Host,
		DialTimeout: 3 * time.Second,
		DB:          r.Db,
		Username:    r.User,
		Password:    r.Pass,
	})

	_, err := rdb.Ping(context.Background()).Result()

	return err

}
