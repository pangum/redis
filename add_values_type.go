package redis

const (
	addValuesTypeLPush addValuesType = 1
	addValuesTypeRPush addValuesType = 2
	addValuesTypeSAdd  addValuesType = 3
)

type addValuesType uint8
