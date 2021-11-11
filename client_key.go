package redis

import (
	`context`
)

func (c *Client) Expire(ctx context.Context, key string, opts ...putOption) (err error) {
	_options := defaultPutOptions()
	for _, opt := range opts {
		opt.applyPut(_options)
	}
	err = c.getClient(_options.options).Expire(ctx, key, _options.expiration).Err()

	return
}

func (c *Client) Del(ctx context.Context, key string, opts ...option) (err error) {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}
	err = c.getClient(_options).Del(ctx, key).Err()

	return
}
