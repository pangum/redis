package redis

const (
	lenTypeLLen  lenType = 1
	lenTypeZCard lenType = 2
)

type lenType uint8
