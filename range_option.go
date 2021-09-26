package redis

type (
	rangeOption interface {
		applyRange(options *rangeOptions)
	}

	rangeOptions struct {
		*options

		start int64
		stop  int64
	}
)

func defaultRangeOptions() *rangeOptions {
	return &rangeOptions{
		options: defaultOptions(),

		start: 1,
		stop:  1,
	}
}
