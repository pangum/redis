package redis

import (
	`context`

	`github.com/go-redis/redis/v8`
)

func (c *Client) HSet(ctx context.Context, key string, opts ...hashOption) (affected int64, err error) {
	_options := defaultHashOptions()
	for _, opt := range opts {
		opt.applyHash(_options)
	}

	values := make([]interface{}, 0, 2*len(_options.fields))
	for _, _field := range _options.fields {
		_format := _field.format
		if formatUnknown == _format {
			_format = _options.options.format
		}

		var marshaled interface{}
		if marshaled, err = c.marshal(_field.value, _format); nil != err {
			return
		}
		values = append(values, _field.key, marshaled)
	}
	affected, err = c.getClient(_options.options).HSet(ctx, key, values...).Result()

	return
}

func (c *Client) HGet(ctx context.Context, key string, field string, value interface{}, opts ...hashOption) (exist bool, err error) {
	_options := defaultHashOptions()
	for _, opt := range opts {
		opt.applyHash(_options)
	}

	var cmd *redis.StringCmd
	defer func() {
		exist = redis.Nil != cmd.Err()
	}()

	if cmd = c.getClient(_options.options).HGet(ctx, key, field); nil != cmd.Err() {
		err = cmd.Err()
	} else {
		err = c.unmarshal(cmd.Val(), value, _options.format)
	}

	return
}

func (c *Client) HIncrBy(ctx context.Context, key string, field string, value int64, opts ...hashOption) (affected int64, err error) {
	_options := defaultHashOptions()
	for _, opt := range opts {
		opt.applyHash(_options)
	}
	affected, err = c.getClient(_options.options).HIncrBy(ctx, key, field, value).Result()

	return
}
