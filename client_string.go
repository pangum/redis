package redis

import (
	`context`

	`github.com/go-redis/redis/v8`
)

func (c *Client) Set(ctx context.Context, key string, value interface{}, opts ...option) (err error) {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	if value, err = c.marshal(value, _options.label, _options.serializer); nil != err {
		return
	}
	err = c.getClient(_options).Set(ctx, key, value, _options.expiration).Err()

	return
}

func (c *Client) Get(ctx context.Context, key string, value interface{}, opts ...option) (exist bool, err error) {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	var cmd *redis.StringCmd
	defer func() {
		exist = redis.Nil != cmd.Err()
	}()

	if cmd = c.getClient(_options).Get(ctx, key); nil != cmd.Err() {
		err = cmd.Err()
	} else {
		err = c.unmarshal(cmd.Val(), value, _options.label, _options.serializer)
	}

	return
}
