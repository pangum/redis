package redis

import (
	`context`
)

func (c *Client) LPush(ctx context.Context, key string, opts ...valuesOption) (affected int64, err error) {
	return c.putValues(ctx, key, putValuesTypeLPush, opts...)
}

func (c *Client) RPush(ctx context.Context, key string, opts ...valuesOption) (affected int64, err error) {
	return c.putValues(ctx, key, putValuesTypeRPush, opts...)
}

func (c *Client) LRange(ctx context.Context, key string, values interface{}, opts ...rangeOption) (err error) {
	return c._range(ctx, key, values, rangeTypeLRange, opts...)
}

func (c *Client) LLen(ctx context.Context, key string, opts ...option) (total int64, err error) {
	return c.len(ctx, key, lenTypeLLen, opts...)
}
