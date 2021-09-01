package redis

var _ sortedSetOption = (*optionValue)(nil)

type optionValue struct {
	value interface{}
}

// Value 配置值
func Value(value interface{}) *optionValue {
	return &optionValue{
		value: value,
	}
}

func (v *optionValue) applySortedSet(options *sortedSetOptions) {
	options.values = append(options.values, v.value)
}
