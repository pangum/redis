package redis

import (
	`context`

	`github.com/go-redis/redis/v8`
	`github.com/storezhang/pangu`
)

func newRedis(config *pangu.Config) (client *Client, err error) {
	panguConfig := new(panguConfig)
	if err = config.Load(panguConfig); nil != err {
		return
	}

	options := &redis.Options{
		Addr:     panguConfig.Redis.Addr,
		Username: panguConfig.Redis.Username,
		Password: panguConfig.Redis.Password,
		DB:       panguConfig.Redis.DB,
	}

	redisClient := redis.NewClient(options)

	if panguConfig.Redis.Ping {
		if _, err = redisClient.Ping(context.Background()).Result(); nil != err {
			return
		}
	}

	client = &Client{Client: redisClient}

	return
}
