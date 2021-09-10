package redis

type field struct {
	key        string
	value      interface{}
	serializer serializer
}
