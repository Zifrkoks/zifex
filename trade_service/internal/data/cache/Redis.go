package cache

import (
	"github.com/redis/go-redis/v9"
)

func SetUpCache() (client *redis.Client) {
	opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
	if err != nil {
		panic("redis dont work")
	}
	client = redis.NewClient(opt)
	return
}
