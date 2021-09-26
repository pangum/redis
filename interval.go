package redis

import (
	`fmt`
)

type interval struct {
	interval string
}

// CloseInterval 闭区间
func CloseInterval(value float64) *interval {
	return &interval{
		interval: fmt.Sprintf("%f", value),
	}
}

// OpenInterval 开区间
func OpenInterval(value float64) *interval {
	return &interval{
		interval: fmt.Sprintf("(%f", value),
	}
}
