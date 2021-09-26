package redis

type (
	countOption interface {
		applyCount(options *countOptions)
	}

	countOptions struct {
		*options

		count      int
		withScores bool
	}
)

func defaultCountOptions() *countOptions {
	return &countOptions{
		options: defaultOptions(),

		count:      1,
		withScores: false,
	}
}
