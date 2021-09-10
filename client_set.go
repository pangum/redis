package redis

import (
	`context`
)

func (c *Client) SAdd(ctx context.Context, key string, opts ...setOption) (affected int64, err error) {
	_options := defaultSetOptions()
	for _, opt := range opts {
		opt.applySet(_options)
	}

	members := make([]interface{}, 0, len(_options.members))
	for _, value := range _options.members {
		var marshaled interface{}
		if marshaled, err = c.marshal(value, _options.label, _options.serializer); nil != err {
			return
		}

		members = append(members, marshaled)
	}
	affected, err = c.getClient(_options.options).SAdd(ctx, key, members...).Result()

	return
}

func (c *Client) SCard(ctx context.Context, key string, opts ...setOption) (total int64, err error) {
	_options := defaultSetOptions()
	for _, opt := range opts {
		opt.applySet(_options)
	}

	if redisCmd := c.getClient(_options.options).SCard(ctx, key); nil != redisCmd.Err() {
		err = redisCmd.Err()
	} else {
		total = redisCmd.Val()
	}

	return
}

func (c *Client) SRem(ctx context.Context, key string, opts ...setOption) (affected int64, err error) {
	_options := defaultSetOptions()
	for _, opt := range opts {
		opt.applySet(_options)
	}

	members := make([]interface{}, 0, len(_options.members))
	for _, _member := range _options.members {
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
