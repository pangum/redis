package redis

import (
	`context`

	`github.com/go-redis/redis/v8`
)

func (c *Client) ZAdd(ctx context.Context, key string, opts ...membersOption) (affected int64, err error) {
	_options := defaultMembersOptions()
	for _, opt := range opts {
		opt.applyMembers(_options)
	}

	zs := make([]*redis.Z, 0, len(_options.members))
	for _, _member := range _options.members {
		var marshaled interface{}
		if marshaled, err = c.marshal(_member.value, _options.label, _options.serializer); nil != err {
			return
		}

		zs = append(zs, &redis.Z{
			Score:  _member.score,
			Member: marshaled,
		})
	}
	affected, err = c.getClient(_options.options).ZAdd(ctx, key, zs...).Result()

	return
}

func (c *Client) ZRange(ctx context.Context, key string, values interface{}, opts ...rangeOption) (err error) {
	_options := defaultRangeOptions()
	for _, opt := range opts {
		opt.applyRange(_options)
	}

	if cmd := c.getClient(_options.options).ZRange(ctx, key, _options.start, _options.stop); nil != cmd.Err() {
		err = cmd.Err()
	} else {
		err = c.unmarshalSlice(cmd.Val(), values, _options.label, _options.serializer)
	}

	return
}

func (c *Client) ZRandMember(ctx context.Context, key string, values interface{}, opts ...countOption) (err error) {
	_options := defaultCountOptions()
	for _, opt := range opts {
		opt.applyCount(_options)
	}

	if cmd := c.getClient(_options.options).ZRandMember(ctx, key, _options.count, _options.withScores); nil != cmd.Err() {
		err = cmd.Err()
	} else {
		err = c.unmarshalSlice(cmd.Val(), values, _options.options.label, _options.options.serializer)
	}

	return
}

func (c *Client) ZCard(ctx context.Context, key string, opts ...option) (total int64, err error) {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	if redisCmd := c.getClient(_options).ZCard(ctx, key); nil != redisCmd.Err() {
		err = redisCmd.Err()
	} else {
		total = redisCmd.Val()
	}

	return
}

func (c *Client) ZCount(ctx context.Context, key string, opts ...intervalOption) (total int64, err error) {
	_options := defaultIntervalOptions()
	for _, opt := range opts {
		opt.applyInterval(_options)
	}

	if redisCmd := c.getClient(_options.options).ZCount(ctx, key, _options.min.interval, _options.max.interval); nil != redisCmd.Err() {
		err = redisCmd.Err()
	} else {
		total = redisCmd.Val()
	}

	return
}

func (c *Client) ZRem(ctx context.Context, key string, opts ...valuesOption) (total int64, err error) {
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
	if redisCmd := c.getClient(_options.options).ZRem(ctx, key, values...); nil != redisCmd.Err() {
		err = redisCmd.Err()
	} else {
		total = redisCmd.Val()
	}

	return
}
