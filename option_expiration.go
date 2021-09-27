package redis

import (
	`time`
)

var (
	_ option         = (*optionExpiration)(nil)
	_ valuesOption   = (*optionExpiration)(nil)
	_ fieldOption    = (*optionExpiration)(nil)
	_ countOption    = (*optionExpiration)(nil)
	_ rangeOption    = (*optionExpiration)(nil)
	_ membersOption  = (*optionExpiration)(nil)
	_ intervalOption = (*optionExpiration)(nil)
)

type optionExpiration struct {
	expiration time.Duration
}

// Expiration 配置过期时间
func Expiration(expiration time.Duration) *optionExpiration {
	return &optionExpiration{
		expiration: expiration,
	}
}

// KeepExpiration 保持过期时间设置
func KeepExpiration() *optionExpiration {
	return &optionExpiration{
		expiration: -1,
	}
}

func (e *optionExpiration) apply(options *options) {
	options.expiration = e.expiration
}

func (e *optionExpiration) applyField(options *fieldOptions) {
	options.expiration = e.expiration
}

func (e *optionExpiration) applyValues(options *valuesOptions) {
	options.expiration = e.expiration
}

func (e *optionExpiration) applyRange(options *rangeOptions) {
	options.expiration = e.expiration
}

func (e *optionExpiration) applyCount(options *countOptions) {
	options.expiration = e.expiration
}

func (e *optionExpiration) applyMembers(options *membersOptions) {
	options.expiration = e.expiration
}

func (e *optionExpiration) applyInterval(options *intervalOptions) {
	options.expiration = e.expiration
}
