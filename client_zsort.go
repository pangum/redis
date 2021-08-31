package redis

import (
	`context`

	`github.com/go-redis/redis/v8`
)

func (c *Client) ZAdd(ctx context.Context, key string, opts ...sortedSetOption) (affected int64, err error) {
	options := defaultSortedSetOptions()
	for _, opt := range opts {
		opt.applySortedSet(options)
	}

	zs := make([]*redis.Z, 0, len(options.members))
	for _, member := range options.members {
		var marshaled interface{}
		if marshaled, err = c.marshal(member.value, options.options); nil != err {
			return
		}

		zs = append(zs, &redis.Z{
			Score:  member.score,
			Member: marshaled,
		})
	}
	affected, err = c.getClient(options.options).ZAdd(ctx, key, zs...).Result()

	return
}

func (c *Client) ZRange(ctx context.Context, key string, values interface{}, opts ...sortedSetOption) (err error) {
	options := defaultSortedSetOptions()
	for _, opt := range opts {
		opt.applySortedSet(options)
	}

	if cmd := c.getClient(options.options).ZRange(ctx, key, options.start, options.stop); nil != cmd.Err() {
		err = cmd.Err()
	} else {
		err = c.unmarshalSlice(cmd.Val(), &values, options.options)
	}

	return
}

func (c *Client) ZRandMember(ctx context.Context, key string, values interface{}, opts ...sortedSetOption) (err error) {
	options := defaultSortedSetOptions()
	for _, opt := range opts {
		opt.applySortedSet(options)
	}

	if cmd := c.getClient(options.options).ZRandMember(ctx, key, options.count, options.withScores); nil != cmd.Err() {
		err = cmd.Err()
	} else {
		err = c.unmarshalSlice(cmd.Val(), &values, options.options)
	}

	return
}

func (c *Client) ZCard(ctx context.Context, key string, opts ...sortedSetOption) (total int64, err error) {
	options := defaultSortedSetOptions()
	for _, opt := range opts {
		opt.applySortedSet(options)
	}

	if redisCmd := c.getClient(options.options).ZCard(ctx, key); nil != redisCmd.Err() {
		err = redisCmd.Err()
	} else {
		total = redisCmd.Val()
	}

	return
}
