package redis

type score interface {
	// Score 设置分数
	Score(score float64)
}
