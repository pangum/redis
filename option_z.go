package redis

var _ membersOption = (*optionZ)(nil)

type optionZ struct {
	score float64
	value interface{}
}

// Z 配置成员
func Z(score int64, value interface{}) *optionZ {
	return &optionZ{
		score: float64(score),
		value: value,
	}
}

func (z *optionZ) applyMembers(options *membersOptions) {
	options.members = append(options.members, &member{
		score: z.score,
		value: z.value,
	})
}
