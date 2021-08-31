package redis

var _ hashOption = (*hashOptionField)(nil)

type hashOptionField struct {
	key   string
	value interface{}
}

// Field 配置字段
func Field(key string, value interface{}) *hashOptionField {
	return &hashOptionField{
		key:   key,
		value: value,
	}
}

func (f *hashOptionField) applyHash(options *hashOptions) {
	options.fields = append(options.fields, &field{
		key:   f.key,
		value: f.value,
	})
}
