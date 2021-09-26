package redis

var _ valuesOption = (*optionValue)(nil)

type optionValue struct {
	value interface{}
}

// Value 配置值
func Value(value interface{}) *optionValue {
	return &optionValue{
		value: value,
	}
}

// Member 配置值
func Member(value interface{}) *optionValue {
	return &optionValue{
		value: value,
	}
}

func (v *optionValue) applyValues(options *valuesOptions) {
	options.values = append(options.values, v.value)
}
