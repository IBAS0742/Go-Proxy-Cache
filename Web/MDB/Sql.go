package MDB

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"go-proxy-cache/Web/MDB/Base"
)

type DB struct {
	db *sql.DB
}

// 表定义
var (
	ProxyTableName      = "proxytable"
	ProxyTableCreateSql = fmt.Sprintf("create table if not exists %s(id varchar(20) primary key not null,group varchar(20) not null,path varchar(256),method tinyint,uri varchar(256),cache tinyint,cachebyuri tinyint,cachewithbody tinyint,parentid varchar(20),unvalid tinyint,record  tinyint)", ProxyTableName)
	CacheTableName      = "cache"
	CacheTableCreateSql = fmt.Sprintf("create table if not exists %s(id varchar(20) primary key not null,blockindex int,content varchar(2048))", CacheTableName)
)

func (Db *DB) InitDb(dbpath string) {
	var err error
	Db.db, err = sql.Open("sqlite3", dbpath)
	if err != nil {
		panic(err)
	}
	Base.ExecuteSqlNotReturn(Db.db, ProxyTableCreateSql, true)
	Base.ExecuteSqlNotReturn(Db.db, CacheTableCreateSql, true)
}
func (Db *DB) ApiGetProxyApi(group string) {

}
