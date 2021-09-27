package redis

type (
	fieldOption interface {
		applyField(options *fieldOptions)
	}

	fieldOptions struct {
		*options

		fields []*field
	}
)

func defaultFieldOptions() *fieldOptions {
	return &fieldOptions{
		options: defaultOptions(),

		fields: make([]*field, 0, 0),
	}
}
