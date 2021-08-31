package redis

import (
	`github.com/go-redis/redis/v8`
	`github.com/storezhang/pangu`
)

func newRedis(config *pangu.Config) (client *Client, err error) {
	panguConfig := new(panguConfig)
	if err = config.Load(panguConfig); nil != err {
		return
	}

	redisConfig := panguConfig.Redis

	// 加载默认连接
	optionsCache := make(map[string]*redis.Options)
	if "" != redisConfig.Addr {
		defaultOptions := &redis.Options{
			Addr:     redisConfig.Addr,
			Username: redisConfig.Options.Username,
			Password: redisConfig.Options.Password,
			DB:       redisConfig.Options.DB,
		}

		optionsCache[defaultLabel] = defaultOptions
	}

	// 加载带标签的服务器
	for _, server := range redisConfig.Servers {
		serverOptions := &redis.Options{
			Addr:     mustString(server.Addr, redisConfig.Addr),
			Username: mustString(server.Options.Username, redisConfig.Options.Username),
			Password: mustString(server.Options.Password, redisConfig.Options.Password),
			DB:       mustInt(server.Options.DB, redisConfig.Options.DB),
		}

		optionsCache[server.Label] = serverOptions
	}

	client = &Client{
		clientCache:  make(map[string]*redis.Client),
		optionsCache: optionsCache,
	}

	return
}
