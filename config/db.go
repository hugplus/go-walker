package config

import "gorm.io/gorm/logger"

type DBCfg struct {
	DSN          string `mapstructure:"dns" json:"dsn" yaml:"dsn"`                                  //连接参数
	Disable      bool   `mapstructure:"disable" json:"disable" yaml:"disable"`                      //是否启用 默认true
	Driver       string `mapstructure:"driver" json:"driver" yaml:"driver"`                         //数据库类型
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         //全局表前缀，单独定义TableName则不生效
	Singular     bool   `mapstructure:"singular" json:"singular" yaml:"singular"`                   //是否开启全局禁用复数，true表示开启
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	MaxLifetime  int    `mapstructure:"max-lifetime" json:"max-lifetime" yaml:"max-lifetime"`       // 链接重置时间（分）
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`                   // 是否开启Gorm全局日志
	LogZap       bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                      // 是否通过zap写入日志文件
	//Tables       []string `mapstructure:"tables" json:"tables" yaml:"tables"`                         // 多库表映射到库
}

func (c *DBCfg) GetLogMode() logger.LogLevel {
	switch c.LogMode {
	case "silent":
		return logger.Silent
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	default:
		return logger.Info

	}
}

func (c *DBCfg) GetMaxIdleConns() int {
	if c.MaxIdleConns < 1 {
		return 10
	}
	return c.MaxIdleConns
}

func (c *DBCfg) GetMaxOpenConns() int {
	if c.MaxOpenConns < 1 {
		return 30
	}
	return c.MaxOpenConns
}

func (c *DBCfg) GetMaxLifetime() int {
	if c.MaxLifetime < 1 {
		return 120
	}
	return c.MaxLifetime
}
