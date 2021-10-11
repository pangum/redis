package redis

import (
	`context`
)

func (c *Client) LPush(ctx context.Context, key string, opts ...valuesOption) (affected int64, err error) {
	return c.addValues(ctx, key, addValuesTypeLPush, opts...)
}

func (c *Client) RPush(ctx context.Context, key string, opts ...valuesOption) (affected int64, err error) {
	return c.addValues(ctx, key, addValuesTypeRPush, opts...)
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
	return c.len(ctx, key, lenTypeLLen, opts...)
}
