package redis

var _ intervalOption = (*optionInterval)(nil)

type optionInterval struct {
	min *interval
	max *interval
}

// Interval 配置个数
func Interval(min *interval, max *interval) *optionInterval {
	return &optionInterval{
		min: min,
		max: max,
	}
}

func (i *optionInterval) applyInterval(options *intervalOptions) {
	options.min = i.min
	options.max = i.max
}
