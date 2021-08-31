package redis

var _ sortedSetOption = (*sortedSetOptionMember)(nil)

type sortedSetOptionMember struct {
	score float64
	value interface{}
}

// Member 配置成员
func Member(score int64, value interface{}) *sortedSetOptionMember {
	return &sortedSetOptionMember{
		score: float64(score),
		value: value,
	}
}

func (m *sortedSetOptionMember) applySortedSet(options *sortedSetOptions) {
	options.members = append(options.members, &member{
		score: m.score,
		value: m.value,
	})
}
