package redis

const (
	serializerUnknown serializer = ""
	serializerJson    serializer = "json"
	serializerProto   serializer = "proto"
	serializerMsgpack serializer = "msgpack"
	serializerXml     serializer = "xml"
	serializerString  serializer = "string"
	serializerBytes   serializer = "bytes"
	serializerBool    serializer = "bool"
	serializerTime    serializer = "time"
	serializerInt     serializer = "int"
	serializerInt64   serializer = "int64"
	serializerUint64  serializer = "uint64"
	serializerFloat32 serializer = "float32"
	serializerFloat64 serializer = "float64"
)

type serializer string
