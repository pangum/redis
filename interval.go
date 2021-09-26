package redis

import (
	`fmt`
)

type interval struct {
	interval string
}

// Close 闭区间
func Close(value float64) *interval {
	return &interval{
		interval: fmt.Sprintf("%f", value),
	}
}

// Open 开区间
func Open(value float64) *interval {
	return &interval{
		interval: fmt.Sprintf("(%f", value),
	}
}
