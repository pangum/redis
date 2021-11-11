package redis

type (
	option interface {
		apply(options *options)
	}

	options struct {
		label      string
		serializer serializer
	}
)

func defaultOptions() *options {
	return &options{
		label:      defaultLabel,
		serializer: serializerJson,
	}
}
