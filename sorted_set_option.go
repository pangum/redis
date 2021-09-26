package redis

type (
	sortedSetOption interface {
		applySortedSet(options *sortedSetOptions)
	}

	sortedSetAddOption interface {
		applySortedSetAdd(options *sortedSetAddOptions)
	}

	sortedSetAddOptions struct {
		*options

		members []*member
	}

	sortedSetOptions struct {
		*options

		values []interface{}

		count      int
		withScores bool

		start int64
		stop  int64
	}
)

func defaultSortedSetAddOptions() *sortedSetAddOptions {
	return &sortedSetAddOptions{
		options: defaultOptions(),

		members: make([]*member, 0, 0),
	}
}

func defaultSortedSetOptions() *sortedSetOptions {
	return &sortedSetOptions{
		options: defaultOptions(),

		values: make([]interface{}, 0, 0),

		count:      1,
		withScores: false,

		start: 1,
		stop:  1,
	}
}
