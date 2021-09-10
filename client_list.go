package redis

import (
	`context`
)

func (c *Client) RPush(ctx context.Context, key string, opts ...listOption) (affected int64, err error) {
	_options := defaultListOptions()
	for _, opt := range opts {
		opt.applyList(_options)
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

func (c *Client) LRange(ctx context.Context, key string, values interface{}, opts ...listOption) (err error) {
	_options := defaultListOptions()
	for _, opt := range opts {
		opt.applyList(_options)
	}

	if cmd := c.getClient(_options.options).LRange(ctx, key, _options.start, _options.stop); nil != cmd.Err() {
		err = cmd.Err()
	} else {
		err = c.unmarshalSlice(cmd.Val(), &values, _options.label, _options.serializer)
	}

	return
}
