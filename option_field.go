package redis

var _ fieldOption = (*optionField)(nil)

type optionField struct {
	key    string
	value  interface{}
	format serializer
}

// Field 配置字段
func Field(key string, value interface{}, opts ...option) *optionField {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	return &optionField{
		key:    key,
		value:  value,
		format: _options.serializer,
	}
}

func (f *optionField) applyField(options *fieldOptions) {
	options.fields = append(options.fields, &field{
		key:        f.key,
		value:      f.value,
		serializer: f.format,
	})
}
