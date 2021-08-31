package redis

var _ sortedSetOption = (*optionCount)(nil)

type optionCount struct {
	count int
}

// Count 配置个数
func Count(count int) *optionCount {
	return &optionCount{
		count: count,
	}
}

func (c *optionCount) applySortedSet(options *sortedSetOptions) {
	options.count = c.count
}
