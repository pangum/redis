package redis

type config struct {
	// 地址
	Addr string `json:"addr" yaml:"addr" xml:"addr" toml:"addr" validate:"required_without=Servers,hostname_port"`
	// 服务器列表
	Servers []server `json:"servers" yaml:"servers" xml:"servers" validate:"required_without=Addr,dive"`
	// 选项
	Options redisOptions `json:"options" yaml:"options" xml:"options" toml:"options"`
}
