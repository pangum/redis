package redis

const (
	formatUnknown format = ""
	formatJson    format = "json"
	formatProto   format = "proto"
	formatMsgpack format = "msgpack"
	formatXml     format = "xml"
	formatString  format = "string"
	formatBytes   format = "bytes"
	formatBool    format = "bool"
	formatTime    format = "time"
	formatInt     format = "int"
	formatInt64   format = "int64"
	formatUint64  format = "uint64"
	formatFloat32 format = "float32"
	formatFloat64 format = "float64"
)

type format string
