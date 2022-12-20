package Server

type TmpCache struct {
	cacheItem  map[string]string
	cacheProxy map[string]ProxyApi
}

func (tc *TmpCache) Save(key, value string) {
	if tc.cacheItem == nil {
		tc.cacheItem = map[string]string{}
	}
	tc.cacheItem[key] = value
}
func (tc *TmpCache) Get(key string) (string, bool) {
	if tc.cacheItem == nil {
		return "", false
	}
	if v, ok := tc.cacheItem[key]; ok {
		return v, true
	}
	return "", false
}

var (
	GetProxyApiAll = "all"
	GetProxyApiOld = "old"
	GetProxyApiNew = "new"
)

func (tc *TmpCache) SaveProxyApi(key string, value ProxyApi) {
	if tc.cacheProxy == nil {
		tc.cacheProxy = map[string]ProxyApi{}
	}
	tc.cacheProxy[key] = value
}
func (tc *TmpCache) GetProxyApi(key string) (ProxyApi, bool) {
	if tc.cacheProxy == nil {
		return ProxyApi{}, false
	}
	if v, ok := tc.cacheProxy[key]; ok {
		return v, true
	}
	return ProxyApi{}, false
}
func (tc TmpCache) GetAllProxyApi() map[string]ProxyApi {
	if tc.cacheProxy != nil {
		return tc.cacheProxy
	} else {
		return map[string]ProxyApi{}
	}
}
func (tc TmpCache) GetOldProxyApi() map[string]ProxyApi {
	if tc.cacheProxy != nil {
		old := map[string]ProxyApi{}
		for k, v := range tc.cacheProxy {
			if !v.New {
				old[k] = v
			}
		}
		return old
	} else {
		return map[string]ProxyApi{}
	}
}
func (tc TmpCache) GetNewProxyApi(parentId string) map[string]ProxyApi {
	if tc.cacheProxy != nil {
		_new := map[string]ProxyApi{}
		for k, v := range tc.cacheProxy {
			if !v.New && v.getId(true) == parentId {
				_new[k] = v
			}
		}
		return _new
	} else {
		return map[string]ProxyApi{}
	}
}
