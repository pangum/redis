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
	clientCache     map[string]*redis.Client
	optionsCache    map[string]*redis.Options
	serializerCache map[string]serializer

	mutex sync.Mutex
}

func (c *Client) Redis(opts ...option) *redis.Client {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	return c.getClient(_options)
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

func (c *Client) marshal(from interface{}, label string, serializer serializer) (to interface{}, err error) {
	serializer = c.getSerializer(label, serializer)
	switch serializer {
	case serializerProto:
		to, err = proto.Marshal(from.(proto.Message))
	case serializerJson:
		to, err = json.Marshal(from)
	case serializerXml:
		to, err = xml.Marshal(from)
	case serializerMsgpack:
		to, err = msgpack.Marshal(from)
	case serializerBytes:
		to = from.([]byte)
	case serializerString:
		to = from.(string)
	case serializerInt:
		to = from.(int)
	case serializerInt64:
		to = from.(int64)
	case serializerUint64:
		to = from.(uint64)
	case serializerBool:
		to = from.(bool)
	case serializerFloat32:
		to = from.(float32)
	case serializerFloat64:
		to = from.(float64)
	case serializerTime:
		to = from.(time.Time)
	}

	return
}

func (c *Client) unmarshal(from string, to interface{}, label string, serializer serializer) (err error) {
	serializer = c.getSerializer(label, serializer)
	switch serializer {
	case serializerProto:
		err = proto.Unmarshal(stringToBytes(from), to.(proto.Message))
	case serializerJson:
		err = json.Unmarshal(stringToBytes(from), to)
	case serializerXml:
		err = xml.Unmarshal(stringToBytes(from), to)
	case serializerMsgpack:
		err = msgpack.Unmarshal(stringToBytes(from), to)
	case serializerBytes:
		toSlice := to.(*[]byte)
		*toSlice = stringToBytes(from)
	case serializerString:
		toString := to.(*string)
		*toString = from
	case serializerInt:
		toInt := to.(*int)
		*toInt, err = strconv.Atoi(from)
	case serializerInt64:
		toInt64 := to.(*int64)
		*toInt64, err = strconv.ParseInt(from, 10, 64)
	case serializerUint64:
		toUint64 := to.(*uint64)
		*toUint64, err = strconv.ParseUint(from, 10, 64)
	case serializerBool:
		toBool := to.(*bool)
		*toBool, err = strconv.ParseBool(from)
	case serializerFloat32:
		toFloat32 := to.(*float32)
		if value, float32Err := strconv.ParseFloat(from, 32); nil == err {
			*toFloat32 = float32(value)
		} else {
			err = float32Err
		}
	case serializerFloat64:
		toFloat64 := to.(*float64)
		*toFloat64, err = strconv.ParseFloat(from, 64)
	case serializerTime:
		to, err = time.Parse(time.RFC3339Nano, from)
	}

	return
}

func (c *Client) unmarshalSlice(strings []string, to interface{}, label string, serializer serializer) (err error) {
	sliceType := reflect.TypeOf(to).Elem()
	elementType := sliceType.Elem()
	newTo := reflect.MakeSlice(sliceType, 0, len(strings))
	for _, str := range strings {
		value := reflect.New(elementType).Interface()
		if err = c.unmarshal(str, value, label, serializer); nil != err {
			return
		}
		newTo = reflect.Append(newTo, reflect.ValueOf(value).Elem())
	}
	reflect.ValueOf(to).Elem().Set(newTo)

	return
}

func (c *Client) getSerializer(label string, original serializer) (serializer serializer) {
	if serializerUnknown == original {
		serializer = c.serializerCache[label]
	} else {
		serializer = original
	}

	return
}
