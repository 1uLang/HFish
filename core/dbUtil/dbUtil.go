package dbUtil

import (
	"HFish/utils/conf"
	"HFish/utils/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose"
	_ "github.com/mattn/go-sqlite3"
)

var engin *gorose.Connection

func init() {
	var err error
	dbType := conf.Get("admin", "db_type")
	dbStr := conf.Get("admin", "db_str")
	dbMaxOpen := conf.GetInt("admin", "db_max_open")
	dbMaxIdle := conf.GetInt("admin", "db_max_idle")
	if dbType == "sqlite" {
		engin, err = gorose.Open(&gorose.DbConfigSingle{Driver: "sqlite3", Dsn: dbStr, SetMaxOpenConns: dbMaxOpen, SetMaxIdleConns: dbMaxIdle})

		if err != nil {
			log.Pr("HFish", "127.0.0.1", "连接 Sqlite 数据库失败", err)
		}
	} else if dbType == "mysql" {
		engin, err = gorose.Open(&gorose.DbConfigSingle{Driver: "mysql", Dsn: dbStr, SetMaxOpenConns: dbMaxOpen, SetMaxIdleConns: dbMaxIdle})

		if err != nil {
			log.Pr("HFish", "127.0.0.1", "连接 Mysql 数据库失败", err)
		}
	}
}

func DB() *gorose.Session {
	return engin.NewSession()
}
