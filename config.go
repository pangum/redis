package redis

type config struct {
	// 服务器地址
	Addr string `json:"addr" yaml:"addr" validate:"required"`
	// 授权，用户名
	Username string `json:"username" yaml:"username"`
	// 授权，密码
	Password string `json:"password" yaml:"password"`
	// 数据库编号
	DB int `json:"db" yaml:"db"`

	// 是否连接时使用Ping测试数据库连接是否完好
	Ping bool `json:"ping" yaml:"ping"`
}
