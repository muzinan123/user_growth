package dbhelper

import (
	"fmt"
	xlog "log"
	"os"
	"time"
	"user_growth/conf"

	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

var dbEngine *xorm.Engine

// InitDb when main function processing, just once process.
func InitDb() {
	if dbEngine != nil {
		return
	}
	// config database source info
	sourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		conf.GlobalConfig.Db.Username,
		conf.GlobalConfig.Db.Password,
		conf.GlobalConfig.Db.Host,
		conf.GlobalConfig.Db.Port,
		conf.GlobalConfig.Db.Database,
		conf.GlobalConfig.Db.Charset)
	if engine, err := xorm.NewEngine(conf.GlobalConfig.Db.Engine, sourceName); err != nil {
		xlog.Fatalf("dbhelper.initDb(%s) error%s\n", sourceName, err.Error())
		return
	} else {
		dbEngine = engine
	}
	// sql log write to stdout
	logger := log.NewSimpleLogger(os.Stdout)
	logger.ShowSQL(conf.GlobalConfig.Db.ShowSql)
	dbEngine.SetLogger(logger)
	if conf.GlobalConfig.Db.ShowSql {
		dbEngine.SetLogLevel(log.DEFAULT_LOG_LEVEL)
	} else {
		dbEngine.SetLogLevel(log.LOG_ERR)
	}
	// more database config
	if conf.GlobalConfig.Db.MaxIdleConns > 0 {
		dbEngine.SetMaxIdleConns(conf.GlobalConfig.Db.MaxIdleConns)
	}
	if conf.GlobalConfig.Db.MaxOpenConns > 0 {
		dbEngine.SetMaxOpenConns(conf.GlobalConfig.Db.MaxOpenConns)
	}
	if conf.GlobalConfig.Db.ConnMaxLifetime > 0 {
		dbEngine.SetConnMaxLifetime(time.Minute * time.Duration(conf.GlobalConfig.Db.ConnMaxLifetime))
	}
}

// GetDb get database engine instance
func GetDb() *xorm.Engine {
	return dbEngine
}
