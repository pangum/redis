package redis

var (
	_ option         = (*optionLabel)(nil)
	_ valuesOption   = (*optionLabel)(nil)
	_ fieldOption    = (*optionLabel)(nil)
	_ rangeOption    = (*optionLabel)(nil)
	_ countOption    = (*optionLabel)(nil)
	_ membersOption  = (*optionLabel)(nil)
	_ intervalOption = (*optionLabel)(nil)
)

type optionLabel struct {
	label string
}

// Label 配置使用哪一个客户端
func Label(label string) *optionLabel {
	return &optionLabel{
		label: label,
	}
}

// Default 配置使用默认客户端
func Default() *optionLabel {
	return &optionLabel{
		label: defaultLabel,
	}
}

func (l *optionLabel) apply(options *options) {
	options.label = l.label
}

func (l *optionLabel) applyValues(options *valuesOptions) {
	options.label = l.label
}

func (l *optionLabel) applyField(options *fieldOptions) {
	options.label = l.label
}

func (l *optionLabel) applyRange(options *rangeOptions) {
	options.label = l.label
}

func (l *optionLabel) applyCount(options *countOptions) {
	options.label = l.label
}

func (l *optionLabel) applyMembers(options *membersOptions) {
	options.label = l.label
}

func (l *optionLabel) applyInterval(options *intervalOptions) {
	options.label = l.label
}
