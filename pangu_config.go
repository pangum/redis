package redis

type panguConfig struct {
	// Redis数据库配置
	Redis config `json:"redis" yaml:"redis" xml:"redis" toml:"redis" validate:"required"`
}
