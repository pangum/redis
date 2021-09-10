package redis

type (
	fieldOptions struct {
		serializer serializer
	}

	fieldOption interface {
		applyField(options *fieldOptions)
	}
)

func defaultFieldOptions() *fieldOptions {
	return &fieldOptions{
		serializer: serializerUnknown,
	}
}
