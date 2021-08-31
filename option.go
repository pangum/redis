package redis

import (
	`time`
)

func defaultOptions() *options {
	return &options{
		label:      defaultLabel,
		format:     formatJson,
		expiration: -1,
	}
}

type (
	option interface {
		apply(options *options)
	}

	options struct {
		label      string
		format     format
		expiration time.Duration
	}
)
