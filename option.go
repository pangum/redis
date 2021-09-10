package redis

import (
	`time`
)

func defaultOptions() *options {
	return &options{
		label:      defaultLabel,
		expiration: -1,
	}
}

type (
	option interface {
		apply(options *options)
	}

	options struct {
		label      string
		serializer serializer
		expiration time.Duration
	}
)
