package Tables

import (
	"database/sql"
	"go-proxy-cache/Web/MDB/DBOP"
)

type CacheBase struct {
	Id         string
	BlockIndex int
	Content    string
}

var CacheTable = DBOP.Table{
	Name: "Cache",
	Fields: []DBOP.TableField{
		{
			Name:       "Id",
			Type:       DBOP.TableFieldTypeCHARACTER36,
			PrimaryKey: true,
			NotNull:    true,
		},
		{
			Name:       "BlockIndex",
			Type:       DBOP.TableFieldTypeInt,
			PrimaryKey: false,
			NotNull:    true,
		},
		{
			Name:       "Content",
			Type:       DBOP.TableFieldTypeCHARACTER36,
			PrimaryKey: false,
			NotNull:    true,
		},
	},
}

func CacheSelectToBase(db *sql.DB, where string) []CacheBase {
	rets := CacheTable.SelectAll(db, where)
	l := len(rets)
	pabs := []CacheBase{}
	for i := 0; i < l; i++ {
		pab := CacheBase{}
		CacheTable.SelectResultToObj(&pab, rets[i])
		pabs = append(pabs, pab)
	}
	return pabs
}
