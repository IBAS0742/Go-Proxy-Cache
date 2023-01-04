package Utils

import (
	"fmt"
	"go-proxy-cache/Web/MDB"
	"go-proxy-cache/Web/MDB/Tables"
	"go-proxy-cache/Web/Server"
)

//type DbProxyApi Tables.ProxyApiBase
//type ServerProxyApi Server.ProxyApi
// 将数据库查询到的 proxy 转化为 server 目录下的 proxyapi
func ChangeProxyApi(api Tables.ProxyApiBase) Server.ProxyApi {
	return Server.ProxyApi{
		Id:            api.Id,
		New:           false,
		Group:         api.Group,
		Path:          api.Path,
		Method:        api.Method,
		Uri:           api.Uri,
		Cache:         api.Cache == 0,
		CacheByURI:    api.CacheByURI == 0,
		CacheWithBody: api.CacheWithBody == 0,
		ParentId:      "",
		//UnValid:       false,
		//Record:        false,
	}
}

// 装载所有的 proxy 接口到缓存钟（启动时，需要执行一次）
func LoadProxyApi(db MDB.DB, server *Server.MainServer) {
	proxys := db.SelectProxyApi("")
	for _, p := range proxys {
		np := ChangeProxyApi(p)
		server.AddProxyApi(np)
		fmt.Println(np)
	}
}
