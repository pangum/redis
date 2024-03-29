package redis

type redisOptions struct {
	// 用户名
	Username string `json:"username" yaml:"username" xml:"username" toml:"username"`
	// 密码
	Password string `json:"password" yaml:"password" xml:"password" toml:"password"`
	// 数据库编号
	DB int `json:"db" yaml:"db" xml:"db" toml:"db"`
	// 序列化器，默认使用Msgpack做序列化
	Serializer serializer `default:"msgpack" json:"serializer" yaml:"serializer" xml:"serializer" toml:"serializer" validate:"oneof=json msgpack proto xml"`
	// 是否连接时使用Ping测试数据库连接是否完好
	Ping bool `default:"true" json:"ping" yaml:"ping" xml:"ping" toml:"ping"`
}
