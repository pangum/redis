package redis

type server struct {
	// 标签
	Label string `json:"label" yaml:"label" xml:"label" toml:"label" validate:"required"`
	// 地址，可以不填，如果不填的话，使用默认地址
	Addr string `json:"addr" yaml:"addr" xml:"addr" toml:"addr" validate:"omitempty,hostname_port"`
	// 选项
	Options redisOptions `json:"options" yaml:"options" xml:"options" toml:"options"`
}
