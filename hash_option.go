package redis

type (
	hashOption interface {
		applyHash(options *hashOptions)
	}

	hashOptions struct {
		*options

		fields []*field
	}
)

func defaultHashOptions() *hashOptions {
	return &hashOptions{
		options: defaultOptions(),

		fields: make([]*field, 0, 0),
	}
}
