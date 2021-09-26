package redis

import (
	`fmt`
)

type interval struct {
	interval string
}

// Close 闭区间
func Close(value int64) *interval {
	return &interval{
		interval: fmt.Sprintf("%d", value),
	}
}

// Open 开区间
func Open(value int64) *interval {
	return &interval{
		interval: fmt.Sprintf("(%d", value),
	}
}
