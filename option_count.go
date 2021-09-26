package redis

var _ countOption = (*optionCount)(nil)

type optionCount struct {
	count      int
	withScores bool
}

// Count 配置个数
func Count(count int) *optionCount {
	return &optionCount{
		count:      count,
		withScores: false,
	}
}

// CountWithScores 配置个数
func CountWithScores(count int) *optionCount {
	return &optionCount{
		count:      count,
		withScores: true,
	}
}

func (c *optionCount) applyCount(options *countOptions) {
	options.count = c.count
	options.withScores = c.withScores
}
