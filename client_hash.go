package redis

import (
	`context`

	`github.com/go-redis/redis/v8`
)

func (c *Client) HSet(ctx context.Context, key string, opts ...hashOption) (affected int64, err error) {
	options := defaultHashOptions()
	for _, opt := range opts {
		opt.applyHash(options)
	}

	values := make([]interface{}, 0, 2*len(options.fields))
	for _, field := range options.fields {
		var marshaled interface{}
		if marshaled, err = c.marshal(field.value, options.options); nil != err {
			return
		}

		values = append(values, field.key, marshaled)
	}
	affected, err = c.getClient(options.options).HSet(ctx, key, values...).Result()

	return
}

func (c *Client) HGet(ctx context.Context, key string, field string, value interface{}, opts ...option) (exist bool, err error) {
	options := defaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}

	var cmd *redis.StringCmd
	defer func() {
		exist = redis.Nil != cmd.Err()
	}()

	if cmd = c.getClient(options).HGet(ctx, key, field); nil != cmd.Err() {
		err = cmd.Err()
	} else {
		err = c.unmarshal(cmd.Val(), value, options)
	}

	return
}
