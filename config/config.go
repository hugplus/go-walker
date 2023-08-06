package config

type AppCfg struct {
	Server ServerCfg         `mapstructure:"server"`
	Logger LogCfg            `mapstructure:"logger"`
	JWT    JWT               `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Dbs    []DBCfg           `mapstructure:"dbs"`
	Cache  RedisCfg          `mapstructure:"cache"`
	Cors   CORS              `mapstructure:"cors" json:"cors" yaml:"cors"`
	Extend map[string]string `mapstructure:"extend"`
}

type ServerCfg struct {
	Mode string `mapstructure:"mode"`
	Host string `mapstructure:"host"`
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
}
