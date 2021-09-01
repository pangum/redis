package redis

type (
	fieldOptions struct {
		format format
	}

	fieldOption interface {
		applyField(options *fieldOptions)
	}
)

func defaultFieldOptions() *fieldOptions {
	return &fieldOptions{
		format: formatUnknown,
	}
}
