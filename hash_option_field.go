package redis

var _ hashOption = (*hashOptionField)(nil)

type hashOptionField struct {
	key    string
	value  interface{}
	format format
}

// Field 配置字段
func Field(key string, value interface{}, opts ...fieldOption) *hashOptionField {
	options := defaultFieldOptions()
	for _, opt := range opts {
		opt.applyField(options)
	}

	return &hashOptionField{
		key:    key,
		value:  value,
		format: options.format,
	}
}

func (f *hashOptionField) applyHash(options *hashOptions) {
	options.fields = append(options.fields, &field{
		key:    f.key,
		value:  f.value,
		format: f.format,
	})
}
