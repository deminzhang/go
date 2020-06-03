package Sql

//ORM
//https://www.kancloud.cn/kancloud/xorm-manual-zh-cn/56004

import (
	"github.com/go-xorm/xorm"
	// _ "github.com/go-sql-driver/mysql"
	// _ "github.com/lib/pq"
	// _ "github.com/mattn/go-sqlite3"
)

var _orm *xorm.Engine

func ORM() *xorm.Engine {
	return _orm
}

//连接/重连
func ORMConnect(driver, source string) *xorm.Engine {
	_driver, _source = driver, source
	var err error
	_orm, err = xorm.NewEngine(driver, source)
	if err != nil {
		panic(err)
	}
	//_orm.ShowSQL(true) //则会在控制台打印出生成的SQL语句；
	_db = _orm.DB().DB
	//缓存
	// cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	// _orm.SetDefaultCacher(cacher)

	return _orm
}
