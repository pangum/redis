package redis

import (
	`context`

	`github.com/go-redis/redis/v8`
)

func (c *Client) Set(ctx context.Context, key string, value interface{}, opts ...option) (err error) {
	options := defaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}

	if value, err = c.marshal(value, options.format); nil != err {
		return
	}
	err = c.getClient(options).Set(ctx, key, value, options.expiration).Err()

	return
}

func (c *Client) Get(ctx context.Context, key string, value interface{}, opts ...option) (exist bool, err error) {
	options := defaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}

	var cmd *redis.StringCmd
	defer func() {
		exist = redis.Nil != cmd.Err()
	}()

	if cmd = c.getClient(options).Get(ctx, key); nil != cmd.Err() {
		err = cmd.Err()
	} else {
		err = c.unmarshal(cmd.Val(), value, options.format)
	}

	return
}
