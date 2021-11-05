package redis

import (
	`context`
)

func (c *Client) SAdd(ctx context.Context, key string, opts ...valuesOption) (affected int64, err error) {
	return c.addValues(ctx, key, addValuesTypeSAdd, opts...)
}

func (c *Client) SCard(ctx context.Context, key string, opts ...option) (total int64, err error) {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	if redisCmd := c.getClient(_options).SCard(ctx, key); nil != redisCmd.Err() {
		err = redisCmd.Err()
	} else {
		total = redisCmd.Val()
	}

	return
}

func (c *Client) SMembers(ctx context.Context, key string, values interface{}, opts ...option) (err error) {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	if cmd := c.getClient(_options).SMembers(ctx, key); nil != cmd.Err() {
		err = cmd.Err()
	} else {
		err = c.unmarshalSlice(cmd.Val(), values, _options.label, _options.serializer)
	}

	return
}

func (c *Client) SRem(ctx context.Context, key string, opts ...valuesOption) (affected int64, err error) {
	_options := defaultValuesOptions()
	for _, opt := range opts {
		opt.applyValues(_options)
	}

	members := make([]interface{}, 0, len(_options.values))
	for _, _member := range _options.values {
		var marshaled interface{}
		if marshaled, err = c.marshal(_member, _options.label, _options.serializer); nil != err {
			return
		}

		members = append(members, marshaled)
	}
	if redisCmd := c.getClient(_options.options).SRem(ctx, key, members...); nil != redisCmd.Err() {
		err = redisCmd.Err()
	} else {
		affected = redisCmd.Val()
	}

	return
}
