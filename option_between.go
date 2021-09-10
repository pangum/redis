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

// Page 模拟分页
func Page(page int64, size int64) *optionBetween {
	return &optionBetween{
		start: (page - 1) * size,
		stop:  page*size - 1,
	}
}

func (b *optionBetween) applySortedSet(options *sortedSetOptions) {
	options.start = b.start
	options.stop = b.stop
}

func (b *optionBetween) applyList(options *listOptions) {
	options.start = b.start
	options.stop = b.stop
}
