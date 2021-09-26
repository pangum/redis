package redis

type (
	valuesOption interface {
		applyValues(options *valuesOptions)
	}

	valuesOptions struct {
		*options

		values []interface{}
	}
)

func defaultValuesOptions() *valuesOptions {
	return &valuesOptions{
		options: defaultOptions(),

		values: make([]interface{}, 0, 0),
	}
}
