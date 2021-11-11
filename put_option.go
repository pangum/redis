package redis

import (
	`time`
)

type (
	putOption interface {
		applyPut(options *putOptions)
	}

	putOptions struct {
		*options

		expiration time.Duration
	}
)

func defaultPutOptions() *putOptions {
	return &putOptions{
		options: defaultOptions(),

		expiration: -1,
	}
}
