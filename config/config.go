package config

type AppCfg struct {
	Server ServerCfg         `mapstructure:"server" json:"server" yaml:"server"`
	Logger LogCfg            `mapstructure:"logger" json:"logger" yaml:"logger"`
	JWT    JWT               `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	DBs    map[string]DBCfg  `mapstructure:"dbs" json:"dbs" yaml:"dbs"`       // 数据库配置
	Cache  RedisCfg          `mapstructure:"cache" json:"cache" yaml:"cache"` // 缓存
	Cors   CORS              `mapstructure:"cors" json:"cors" yaml:"cors"`
	Extend map[string]string `mapstructure:"extend" json:"extend" yaml:"extend"`
}

type ServerCfg struct {
	Mode         string `mapstructure:"mode" json:"mode" yaml:"mode"`
	Host         string `mapstructure:"host" json:"host" yaml:"host"`
	Name         string `mapstructure:"name" json:"name" yaml:"name"`
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`
	ReadTimeout  int    `mapstructure:"read-timeout" json:"read-timeout" yaml:"read-timeout"`
	WriteTimeout int    `mapstructure:"write-timeout" json:"write-timeout" yaml:"write-timeout"`
}

func (e *ServerCfg) GetReadTimeout() int {
	if e.ReadTimeout < 1 {
		return 10
	}
	return e.ReadTimeout
}

func (e *ServerCfg) GetWriteTimeout() int {
	if e.WriteTimeout < 1 {
		return 10
	}
	return e.WriteTimeout
}
