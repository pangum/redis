package redis

type (
	sortedSetOption interface {
		applySortedSet(options *sortedSetOptions)
	}

	sortedSetOptions struct {
		*options

		members []*member

		count      int
		withScores bool

		start int64
		stop  int64
	}
)

func defaultSortedSetOptions() *sortedSetOptions {
	return &sortedSetOptions{
		options: defaultOptions(),

		members: make([]*member, 0, 0),

		count:      1,
		withScores: true,

		start: 1,
		stop:  1,
	}
}
