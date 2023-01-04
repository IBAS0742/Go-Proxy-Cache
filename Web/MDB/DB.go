package MDB

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"go-proxy-cache/Web/MDB/Base"
	"go-proxy-cache/Web/MDB/Tables"
)

type DB struct {
	db *sql.DB
}

func (Db *DB) InitDb(dbpath string) {
	var err error
	Db.db, err = sql.Open("sqlite3", dbpath)
	if err != nil {
		panic(err)
	}
	Base.ExecuteSqlNotReturn(Db.db, Tables.ProxyApiTable.BuildCreateSql(), true)
	Base.ExecuteSqlNotReturn(Db.db, Tables.CacheTable.BuildCreateSql(), true)
}

// 统一将所有的方法注册到这里
// 注册接口
func (Db DB) SelectProxyApi(where string) []Tables.ProxyApiBase {
	return Tables.ProxyApiSelectToBase(Db.db, where)
}
func (Db DB) SelectCache(where string) []Tables.CacheBase {
	return Tables.CacheSelectToBase(Db.db, where)
}
