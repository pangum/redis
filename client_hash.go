package redis

import (
	`context`

	`github.com/go-redis/redis/v8`
)

func (c *Client) HSet(ctx context.Context, key string, opts ...fieldOption) (affected int64, err error) {
	_options := defaultFieldOptions()
	for _, opt := range opts {
		opt.applyField(_options)
	}

	values := make([]interface{}, 0, 2*len(_options.fields))
	for _, _field := range _options.fields {
		_serializer := _field.serializer
		if serializerUnknown == _serializer {
			_serializer = _options.serializer
		}

		var marshaled interface{}
		if marshaled, err = c.marshal(_field.value, _options.label, _serializer); nil != err {
			return
		}
		values = append(values, _field.key, marshaled)
	}
	affected, err = c.getClient(_options.options).HSet(ctx, key, values...).Result()

	return
}

func (c *Client) HGet(ctx context.Context, key string, field string, value interface{}, opts ...option) (exist bool, err error) {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	var cmd *redis.StringCmd
	defer func() {
		exist = redis.Nil != cmd.Err()
	}()

	if cmd = c.getClient(_options).HGet(ctx, key, field); nil != cmd.Err() {
		err = cmd.Err()
	} else {
		err = c.unmarshal(cmd.Val(), value, _options.label, _options.serializer)
	}

	return
}

func (c *Client) HKeys(ctx context.Context, key string, opts ...option) (keys []string, err error) {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	return c.getClient(_options).HKeys(ctx, key).Result()
}

func (c *Client) HIncrBy(ctx context.Context, key string, field string, value int64, opts ...option) (int64, error) {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	return c.getClient(_options).HIncrBy(ctx, key, field, value).Result()
}

func (c *Client) HDel(ctx context.Context, key string, field string, opts ...option) (int64, error) {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	return c.getClient(_options).HDel(ctx, key, field).Result()
}
