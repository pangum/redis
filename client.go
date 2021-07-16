package redis

import (
	`context`
	`encoding/json`
	`time`

	`github.com/go-redis/redis/v8`
	`github.com/golang/protobuf/proto`
)

type Client struct {
	*redis.Client
}

func (c *Client) Set(key string, value interface{}, expirations ...time.Duration) *redis.StatusCmd {
	return c.Client.Set(context.Background(), key, value, getExpiration(expirations...))
}

func (c *Client) SetJSON(key string, value interface{}, expirations ...time.Duration) *redis.StatusCmd {
	if jsonValue, err := json.Marshal(value); nil == err {
		return c.Client.Set(context.Background(), key, jsonValue, getExpiration(expirations...))
	}

	return nil
}

func (c *Client) SetProtobuf(key string, value proto.Message, expirations ...time.Duration) *redis.StatusCmd {
	if protobuf, err := proto.Marshal(value); nil == err {
		return c.Client.Set(context.Background(), key, protobuf, getExpiration(expirations...))
	}

	return nil
}

func (c *Client) Del(key string) *redis.IntCmd {
	return c.Client.Del(context.Background(), key)
}

func (c *Client) Get(key string) string {
	return c.Client.Get(context.Background(), key).Val()
}

func (c *Client) IncrBy(key string, value int64, expirations ...time.Duration) int64 {
	c.Client.Expire(context.Background(), key, getExpiration(expirations...))

	return c.Client.IncrBy(context.Background(), key, value).Val()
}

func getExpiration(expirations ...time.Duration) time.Duration {
	if 0 != len(expirations) {
		return expirations[0]
	}

	return 0
}
