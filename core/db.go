package core

import (
	"database/sql"
	"errors"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func dbInit() {

	var db *gorm.DB
	var err error

	for _, dbc := range Cfg.Dbs {
		switch dbc.Driver {
		case "mysql":
			db, err = gorm.Open(mysql.Open(dbc.Dsn), &gorm.Config{})
		case "pgsql":
			db, err = gorm.Open(postgres.Open(dbc.Dsn), &gorm.Config{})
		case "sqlite":
			db, err = gorm.Open(sqlite.Open(dbc.Dsn), &gorm.Config{})
		case "mssql":
			db, err = gorm.Open(sqlserver.Open(dbc.Dsn), &gorm.Config{})
		// case "oracle":
		// 	db, err = gorm.Open(oracle.Open(dbc.Dsn), &gorm.Config{})
		// case "clickhouse":
		// 	db, err = gorm.Open(clickhouse.Open(dbc.Dsn), &gorm.Config{})
		default:
			err = errors.New("db err")
		}
		if err != nil {
			panic(err)
		}
		var sqlDB *sql.DB
		sqlDB, err = db.DB()
		if err != nil {
			panic(err)
		}
		sqlDB.SetMaxIdleConns(dbc.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbc.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(dbc.MaxLifetime))
		SetDb(dbc.DBName, db)
	}

}

func SetDb(key string, db *gorm.DB) {
	lock.Lock()
	defer lock.Unlock()
	dbs[key] = db
}

// GetDb 获取所有map里的db数据
func Dbs() map[string]*gorm.DB {
	// lock.RLock()
	// defer lock.RUnlock()
	return dbs
}

func Db(name string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	if db, ok := dbs[name]; !ok || db == nil {
		panic("db not init")
	} else {
		return db
	}
}
