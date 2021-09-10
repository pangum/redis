package redis

type (
	setOptions struct {
		*options

		members []interface{}
	}

	setOption interface {
		applySet(options *setOptions)
	}
)

func defaultSetOptions() *setOptions {
	return &setOptions{
		options: defaultOptions(),

		members: make([]interface{}, 0, 0),
	}
}
