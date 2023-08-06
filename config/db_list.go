package config

type DsnProvider interface {
	Dsn() string
}

// Embeded 结构体可以压平到上一层，从而保持 config 文件的结构和原来一样
// 见 playground: https://go.dev/play/p/KIcuhqEoxmY

type DBConn struct {
	Dsn          string `mapstructure:"dns"`                                                        //连接参数
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         //全局表前缀，单独定义TableName则不生效
	Singular     bool   `mapstructure:"singular" json:"singular" yaml:"singular"`                   //是否开启全局禁用复数，true表示开启
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	MaxLifetime  int    `mapstructure:"max-lifetime" json:"max-lifetime" yaml:"max-lifetime"`       // 链接重置时间（分）
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`                   // 是否开启Gorm全局日志
	LogZap       bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                      // 是否通过zap写入日志文件
}

type DBCfg struct {
	Disable bool   `mapstructure:"disable" json:"disable" yaml:"disable"`
	Driver  string `mapstructure:"driver"` //数据库类型
	DBName  string `mapstructure:"dbname" json:"alias-name" yaml:"alias-name"`
	DBConn  `yaml:",inline" mapstructure:",squash"`
}
