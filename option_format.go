package redis

var _ option = (*optionFormat)(nil)

type optionFormat struct {
	format format
}

// Proto 谷歌Protocol Buffer序列化
func Proto() *optionFormat {
	return &optionFormat{
		format: formatProto,
	}
}

// JSON 使用JSON序列化
func JSON() *optionFormat {
	return &optionFormat{
		format: formatJson,
	}
}

// XML 使用XML序列化
func XML() *optionFormat {
	return &optionFormat{
		format: formatXml,
	}
}

// Msgpack 使用Msgpack序列化
func Msgpack() *optionFormat {
	return &optionFormat{
		format: formatMsgpack,
	}
}

// Bytes 原始数据
func Bytes() *optionFormat {
	return &optionFormat{
		format: formatBytes,
	}
}

// String 字符串数据
func String() *optionFormat {
	return &optionFormat{
		format: formatString,
	}
}

// Int 整形数据
func Int() *optionFormat {
	return &optionFormat{
		format: formatInt,
	}
}

// Int64 整形数据
func Int64() *optionFormat {
	return &optionFormat{
		format: formatInt64,
	}
}

// Uint64 整形数据
func Uint64() *optionFormat {
	return &optionFormat{
		format: formatUint64,
	}
}

// Float32 浮点数据
func Float32() *optionFormat {
	return &optionFormat{
		format: formatFloat32,
	}
}

// Float64 浮点数据
func Float64() *optionFormat {
	return &optionFormat{
		format: formatFloat64,
	}
}

// Bool 布尔数据
func Bool() *optionFormat {
	return &optionFormat{
		format: formatBool,
	}
}

// Time 时间数据
func Time() *optionFormat {
	return &optionFormat{
		format: formatTime,
	}
}

func (f *optionFormat) apply(options *options) {
	options.format = f.format
}

func (f *optionFormat) applySortedSet(options *sortedSetOptions) {
	options.format = f.format
}

func (f *optionFormat) applyHash(options *hashOptions) {
	options.format = f.format
}

func (f *optionFormat) applyField(options *fieldOptions) {
	options.format = f.format
}
