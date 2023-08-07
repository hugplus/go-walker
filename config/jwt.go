package config

type JWT struct {
	SignKey string `mapstructure:"sign-key" json:"sign-key" yaml:"sign-key"` // jwt签名
	Expires int    `mapstructure:"expires" json:"expires" yaml:"expires"`    // 有效时长 分钟
	Refresh int    `mapstructure:"refresh" json:"refresh" yaml:"refresh"`    // 刷新时长
}
