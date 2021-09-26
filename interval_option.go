package redis

type (
	intervalOption interface {
		applyInterval(options *intervalOptions)
	}

	intervalOptions struct {
		*options

		min *interval
		max *interval
	}
)

func defaultIntervalOptions() *intervalOptions {
	return &intervalOptions{
		options: defaultOptions(),

		min: CloseInterval(1),
		max: CloseInterval(1),
	}
}
