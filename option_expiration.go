package redis

import (
	`time`
)

var _ putOption = (*optionExpiration)(nil)

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

func (e *optionExpiration) applyPut(options *putOptions) {
	options.expiration = e.expiration
}
