package MDB

//import (
//	"database/sql"
//	"fmt"
//	"go-proxy-cache/Web/Server"
//)
//
//type Int int
//
//func (i Int) ToBool() bool {
//	if i == 0 {
//		return false
//	} else {
//		return true
//	}
//}
//func BoolTo(b bool) (i Int) {
//	if b {
//		i = 1
//	} else {
//		i = 0
//	}
//	return i
//}
//
//var ProxySelectSql = fmt.Sprintf("select (id,group,path,method,uri,cache,cachebyuri,cachewithbody,parentid,unvalid,record) from %s", ProxyTableName)
//
//type ProxyApiSelectModel struct {
//	Id            string
//	Group         string
//	Path          string
//	Method        int
//	Uri           string
//	Cache         int
//	CacheByUri    int
//	CacheWithBody int
//	ParentId      string
//	Unvalid       int
//	Record        int
//}
//
//func (proxy ProxyApiSelectModel) GetProxyFromSqlResult(rows *sql.Rows) []Server.ProxyApi {
//	apis := []Server.ProxyApi{}
//	for rows.Next() {
//		rows.Scan(
//			&proxy.Id,
//			&proxy.Group,
//			&proxy.Path,
//			&proxy.Method,
//			&proxy.Uri,
//			&proxy.Cache,
//			&proxy.CacheByUri,
//			&proxy.CacheWithBody,
//			&proxy.ParentId,
//			&proxy.Unvalid,
//			&proxy.Record,
//		)
//		api := Server.ProxyApi{
//			Id:    proxy.Id,
//			New:   false,
//			Group: proxy.Group,
//			Path:  proxy.Path,
//			//Method:        proxy.Method,
//			Uri:           proxy.Uri,
//			Cache:         false,
//			CacheByURI:    false,
//			CacheWithBody: false,
//			ParentId:      proxy.ParentId,
//			UnValid:       false,
//			Record:        false,
//		}
//
//		apis = append(apis, api)
//	}
//	return apis
//}
