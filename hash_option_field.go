package redis

var _ hashOption = (*hashOptionField)(nil)

type hashOptionField struct {
	key    string
	value  interface{}
	format serializer
}

// Field 配置字段
func Field(key string, value interface{}, opts ...fieldOption) *hashOptionField {
	_options := defaultFieldOptions()
	for _, opt := range opts {
		opt.applyField(_options)
	}

	return &hashOptionField{
		key:    key,
		value:  value,
		format: _options.serializer,
	}
}

func (f *hashOptionField) applyHash(options *hashOptions) {
	options.fields = append(options.fields, &field{
		key:        f.key,
		value:      f.value,
		serializer: f.format,
	})
}
