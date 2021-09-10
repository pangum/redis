package redis

var (
	_ sortedSetOption = (*optionValue)(nil)
	_ setOption       = (*optionValue)(nil)
	_ listOption      = (*optionValue)(nil)
)

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

func (v *optionValue) applySortedSet(options *sortedSetOptions) {
	options.values = append(options.values, v.value)
}

func (v *optionValue) applySet(options *setOptions) {
	options.members = append(options.members, v.value)
}

func (v *optionValue) applyList(options *listOptions) {
	options.values = append(options.values, v.value)
}
