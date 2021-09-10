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
	serializerCache := make(map[string]serializer)
	if "" != redisConfig.Addr {
		defaultOptions := &redis.Options{
			Addr:     redisConfig.Addr,
			Username: redisConfig.Options.Username,
			Password: redisConfig.Options.Password,
			DB:       redisConfig.Options.DB,
		}

		optionsCache[defaultLabel] = defaultOptions
		serializerCache[defaultLabel] = redisConfig.Options.Serializer
	}

	// 加载带标签的服务器
	for _, _server := range redisConfig.Servers {
		serverOptions := &redis.Options{
			Addr:     mustString(_server.Addr, redisConfig.Addr),
			Username: mustString(_server.Options.Username, redisConfig.Options.Username),
			Password: mustString(_server.Options.Password, redisConfig.Options.Password),
			DB:       mustInt(_server.Options.DB, redisConfig.Options.DB),
		}

		optionsCache[_server.Label] = serverOptions
		serializerCache[_server.Label] = _server.Options.Serializer
	}

	client = &Client{
		clientCache:     make(map[string]*redis.Client),
		optionsCache:    optionsCache,
		serializerCache: serializerCache,
	}

	return
}
