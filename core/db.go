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

	for key, dbc := range Cfg.DBs {
		if !dbc.Disable {

			curLog := NewGormZapLog(Log, dbc.GetLogMode(), 200)
			// logger.New(
			// 	NewGormLogger(), // io writer
			// 	logger.Config{
			// 		SlowThreshold:             time.Second,      // Slow SQL threshold
			// 		LogLevel:                  dbc.GetLogMode(), // Log level
			// 		IgnoreRecordNotFoundError: true,             // Ignore ErrRecordNotFound error for logger
			// 		ParameterizedQueries:      true,             // Don't include params in the SQL log
			// 		Colorful:                  false,            // Disable color
			// 	},
			// )

			switch dbc.Driver {
			case "mysql":
				db, err = gorm.Open(mysql.Open(dbc.DSN), &gorm.Config{
					Logger: curLog,
				})
			case "pgsql":
				db, err = gorm.Open(postgres.Open(dbc.DSN), &gorm.Config{})
			case "sqlite":
				db, err = gorm.Open(sqlite.Open(dbc.DSN), &gorm.Config{})
			case "mssql":
				db, err = gorm.Open(sqlserver.Open(dbc.DSN), &gorm.Config{})
			// case "oracle":
			// 	db, err = gorm.Open(oracle.Open(dbc.DSN), &gorm.Config{})
			// case "clickhouse":
			// 	db, err = gorm.Open(clickhouse.Open(dbc.DSN), &gorm.Config{})
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
			sqlDB.SetMaxIdleConns(dbc.GetMaxIdleConns())
			sqlDB.SetMaxOpenConns(dbc.GetMaxOpenConns())
			sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(dbc.GetMaxLifetime()))
			SetDb(key, db)
		}
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
		// Log.Error("DB not found", zap.Error(errors.New(name+" DB not found")))
		// return nil
	} else {
		return db
	}
}

func DB() *gorm.DB {
	return Db("master")
}
