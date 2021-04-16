package db

import (
	gosql "database/sql"

	"entgo.io/ent/dialect/sql"
	"github.com/gogf/gf/frame/g"
)

func getDriver(master bool) (drv *sql.Driver, err error) {
	var sqlDB *gosql.DB
	if master {
		sqlDB, err = g.DB().Master()
	} else {
		sqlDB, err = g.DB().Slave()
	}
	if err != nil {
		return nil, err
	}

	drv = sql.OpenDB(g.Cfg().GetString("database.default.name"), sqlDB)
	return
}

// GetMasterDriver 主数据库驱动
func GetMasterDriver() (drv *sql.Driver, err error) {
	return getDriver(true)
}

// GetSlaveDriver 从数据库驱动
func GetSlaveDriver() (drv *sql.Driver, err error) {
	return getDriver(false)
}
