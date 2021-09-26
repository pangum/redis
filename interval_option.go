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

		min: Close(1),
		max: Close(1),
	}
}
