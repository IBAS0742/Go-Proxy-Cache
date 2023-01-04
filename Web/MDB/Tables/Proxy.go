package Tables

import (
	"database/sql"
	"go-proxy-cache/Web/MDB/DBOP"
)

type ProxyApiBase struct {
	Id            string // 数据库 id
	New           bool   // 是否是临时新增的（true表示不是从数据库获取的）
	Group         string // 分组(可以用于区分多个使用相同 ip端口 的项目，例如同样都是 http://localhost:8080 的不同项目)
	Uri           string // 请求前缀 (http://localhost:81)
	Path          string // 请求路径（ /api/user/login ）
	Method        string // 请求方法(Post Get Put ...)
	Cache         int    // 是否进行缓存
	CacheByURI    int    // 是否使用 uri 内容标记多个缓存
	CacheWithBody int    // 是否使用 body 内容的 hash 值标记多个缓存
	ParentId      string // 标记父级（用于 new = true 时，标记为捕获的可以被记录的请求）
}

//type ProxyApiSQL struct {
//	Id string
//	Group string
//	Uri string
//	Path string
//	Method string
//	Cache bool
//	CacheByURI bool
//	CacheWithBody bool
//}
var ProxyApiTable = DBOP.Table{
	Name: "Proxy",
	Fields: []DBOP.TableField{
		{
			Name:       "Id",
			Type:       DBOP.TableFieldTypeCHARACTER36,
			PrimaryKey: true,
			NotNull:    true,
		},
		{
			Name:    "Group",
			Type:    DBOP.TableFieldTypeCHARACTER36,
			NotNull: true,
		},
		{
			Name:    "Uri",
			Type:    DBOP.TableFieldTypeCHARACTER256,
			NotNull: true,
		},
		{
			Name:    "Path",
			Type:    DBOP.TableFieldTypeCHARACTER256,
			NotNull: true,
		},
		{
			Name:    "Method",
			Type:    DBOP.TableFieldTypeCHARACTER36,
			NotNull: true,
		},
		{
			Name:    "Cache",
			Type:    DBOP.TableFieldTypeInt,
			NotNull: true,
		},
		{
			Name:    "CacheByURI",
			Type:    DBOP.TableFieldTypeInt,
			NotNull: true,
		},
		{
			Name:    "CacheWithBody",
			Type:    DBOP.TableFieldTypeInt,
			NotNull: true,
		},
	},
}

func ProxyApiSelectToBase(db *sql.DB, where string) []ProxyApiBase {
	rets := ProxyApiTable.SelectAll(db, where)
	l := len(rets)
	pabs := []ProxyApiBase{}
	for i := 0; i < l; i++ {
		pab := ProxyApiBase{}
		ProxyApiTable.SelectResultToObj(&pab, rets[i])
		pabs = append(pabs, pab)
	}
	return pabs
}
