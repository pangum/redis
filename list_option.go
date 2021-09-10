package redis

type (
	listOptions struct {
		*options

		values []interface{}

		count int
		start int64
		stop  int64
	}

	listOption interface {
		applyList(options *listOptions)
	}
)

func defaultListOptions() *listOptions {
	return &listOptions{
		options: defaultOptions(),

		values: make([]interface{}, 0, 0),
	}
}
