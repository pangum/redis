package redis

var (
	_ option         = (*optionSerializer)(nil)
	_ valuesOption   = (*optionSerializer)(nil)
	_ fieldOption    = (*optionSerializer)(nil)
	_ countOption    = (*optionSerializer)(nil)
	_ rangeOption    = (*optionSerializer)(nil)
	_ membersOption  = (*optionSerializer)(nil)
	_ intervalOption = (*optionSerializer)(nil)
)

type optionSerializer struct {
	serializer serializer
}

// Proto 谷歌Protocol Buffer序列化
func Proto() *optionSerializer {
	return &optionSerializer{
		serializer: serializerProto,
	}
}

// JSON 使用JSON序列化
func JSON() *optionSerializer {
	return &optionSerializer{
		serializer: serializerJson,
	}
}

// XML 使用XML序列化
func XML() *optionSerializer {
	return &optionSerializer{
		serializer: serializerXml,
	}
}

// Msgpack 使用Msgpack序列化
func Msgpack() *optionSerializer {
	return &optionSerializer{
		serializer: serializerMsgpack,
	}
}

// Bytes 原始数据
func Bytes() *optionSerializer {
	return &optionSerializer{
		serializer: serializerBytes,
	}
}

// String 字符串数据
func String() *optionSerializer {
	return &optionSerializer{
		serializer: serializerString,
	}
}

// Int 整形数据
func Int() *optionSerializer {
	return &optionSerializer{
		serializer: serializerInt,
	}
}

// Int64 整形数据
func Int64() *optionSerializer {
	return &optionSerializer{
		serializer: serializerInt64,
	}
}

// Uint64 整形数据
func Uint64() *optionSerializer {
	return &optionSerializer{
		serializer: serializerUint64,
	}
}

// Float32 浮点数据
func Float32() *optionSerializer {
	return &optionSerializer{
		serializer: serializerFloat32,
	}
}

// Float64 浮点数据
func Float64() *optionSerializer {
	return &optionSerializer{
		serializer: serializerFloat64,
	}
}

// Bool 布尔数据
func Bool() *optionSerializer {
	return &optionSerializer{
		serializer: serializerBool,
	}
}

// Time 时间数据
func Time() *optionSerializer {
	return &optionSerializer{
		serializer: serializerTime,
	}
}

func (s *optionSerializer) apply(options *options) {
	options.serializer = s.serializer
}

func (s *optionSerializer) applyField(options *fieldOptions) {
	options.serializer = s.serializer
}

func (s *optionSerializer) applyValues(options *valuesOptions) {
	options.serializer = s.serializer
}

func (s *optionSerializer) applyRange(options *rangeOptions) {
	options.serializer = s.serializer
}

func (s *optionSerializer) applyCount(options *countOptions) {
	options.serializer = s.serializer
}

func (s *optionSerializer) applyMembers(options *membersOptions) {
	options.serializer = s.serializer
}

func (s *optionSerializer) applyInterval(options *intervalOptions) {
	options.serializer = s.serializer
}
