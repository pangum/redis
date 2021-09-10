package redis

var _ sortedSetAddOption = (*sortedSetOptionZ)(nil)

type sortedSetOptionZ struct {
	score float64
	value interface{}
}

// Z 配置成员
func Z(score int64, value interface{}) *sortedSetOptionZ {
	return &sortedSetOptionZ{
		score: float64(score),
		value: value,
	}
}

func (z *sortedSetOptionZ) applySortedSetAdd(options *sortedSetAddOptions) {
	options.members = append(options.members, &member{
		score: z.score,
		value: z.value,
	})
}
