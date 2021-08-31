package redis

var _ sortedSetOption = (*optionBetween)(nil)

type optionBetween struct {
	start int64
	stop  int64
}

// Between 配置起止段
func Between(start int64, stop int64) *optionBetween {
	return &optionBetween{
		start: start,
		stop:  stop,
	}
}

func (b *optionBetween) applySortedSet(options *sortedSetOptions) {
	options.start = b.start
}
