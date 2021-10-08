package redis

import (
	`context`
)

func (c *Client) RPush(ctx context.Context, key string, opts ...valuesOption) (affected int64, err error) {
	_options := defaultValuesOptions()
	for _, opt := range opts {
		opt.applyValues(_options)
	}

	values := make([]interface{}, 0, len(_options.values))
	for _, value := range _options.values {
		var marshaled interface{}
		if marshaled, err = c.marshal(value, _options.label, _options.serializer); nil != err {
			return
		}

		values = append(values, marshaled)
	}
	affected, err = c.getClient(_options.options).RPush(ctx, key, values...).Result()

	return
}

func (c *Client) LRange(ctx context.Context, key string, values interface{}, opts ...rangeOption) (err error) {
	_options := defaultRangeOptions()
	for _, opt := range opts {
		opt.applyRange(_options)
	}

	if cmd := c.getClient(_options.options).LRange(ctx, key, _options.start, _options.stop); nil != cmd.Err() {
		err = cmd.Err()
	} else {
		err = c.unmarshalSlice(cmd.Val(), values, _options.label, _options.serializer)
	}

	return
}

func (c *Client) LLen(ctx context.Context, key string, opts ...option) (total int64, err error) {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	if redisCmd := c.getClient(_options).LLen(ctx, key); nil != redisCmd.Err() {
		err = redisCmd.Err()
	} else {
		total = redisCmd.Val()
	}

	return
}
