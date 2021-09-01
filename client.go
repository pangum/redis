package redis

import (
	`encoding/json`
	`encoding/xml`
	`reflect`
	`strconv`
	`sync`
	`time`

	`github.com/go-redis/redis/v8`
	`github.com/golang/protobuf/proto`
	`github.com/vmihailenco/msgpack/v5`
)

// Client Redis客户端
type Client struct {
	clientCache  map[string]*redis.Client
	optionsCache map[string]*redis.Options

	mutex sync.Mutex
}

func (c *Client) Redis(opts ...option) *redis.Client {
	options := defaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}

	return c.getClient(options)
}

func (c *Client) getClient(options *options) (client *redis.Client) {
	c.mutex.Lock()
	defer func() {
		options.label = defaultLabel
		c.mutex.Unlock()
	}()

	var exist bool
	if client, exist = c.clientCache[options.label]; exist {
		return
	}

	client = redis.NewClient(c.optionsCache[options.label])
	c.clientCache[options.label] = client

	return
}

func (c *Client) marshal(from interface{}, format format) (to interface{}, err error) {
	switch format {
	case formatProto:
		to, err = proto.Marshal(from.(proto.Message))
	case formatJson:
		to, err = json.Marshal(from)
	case formatXml:
		to, err = xml.Marshal(from)
	case formatMsgpack:
		to, err = msgpack.Marshal(from)
	case formatBytes:
		to = from.([]byte)
	case formatString:
		to = from.(string)
	case formatInt:
		to = from.(int)
	case formatInt64:
		to = from.(int64)
	case formatUint64:
		to = from.(uint64)
	case formatBool:
		to = from.(bool)
	case formatFloat32:
		to = from.(float32)
	case formatFloat64:
		to = from.(float64)
	case formatTime:
		to = from.(time.Time)
	}

	return
}

func (c *Client) unmarshal(from string, to interface{}, format format) (err error) {
	switch format {
	case formatProto:
		err = proto.Unmarshal(stringToBytes(from), to.(proto.Message))
	case formatJson:
		err = json.Unmarshal(stringToBytes(from), to)
	case formatXml:
		err = xml.Unmarshal(stringToBytes(from), to)
	case formatMsgpack:
		err = msgpack.Unmarshal(stringToBytes(from), to)
	case formatBytes:
		to = stringToBytes(from)
	case formatString:
		to = from
	case formatInt:
		to, err = strconv.Atoi(from)
	case formatInt64:
		to, err = strconv.ParseInt(from, 10, 64)
	case formatUint64:
		to, err = strconv.ParseUint(from, 10, 64)
	case formatBool:
		to, err = strconv.ParseBool(from)
	case formatFloat32:
		to, err = strconv.ParseFloat(from, 32)
	case formatFloat64:
		to, err = strconv.ParseFloat(from, 64)
	case formatTime:
		to, err = time.Parse(time.RFC3339Nano, from)
	}

	return
}

func (c *Client) unmarshalSlice(strings []string, to interface{}, format format) (err error) {
	sliceType := reflect.TypeOf(to).Elem()
	elementType := sliceType.Elem()
	newTo := reflect.MakeSlice(sliceType, 0, len(strings))
	for _, str := range strings {
		value := reflect.New(elementType).Interface()
		if err = c.unmarshal(str, value, format); nil != err {
			return
		}
		newTo = reflect.Append(newTo, reflect.ValueOf(value).Elem())
	}
	reflect.ValueOf(to).Elem().Set(newTo)

	return
}
