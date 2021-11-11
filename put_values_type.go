package redis

const (
	putValuesTypeLPush putValuesType = 1
	putValuesTypeRPush putValuesType = 2
	putValuesTypeSAdd  putValuesType = 3
)

type putValuesType uint8
