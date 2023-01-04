package Server

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type Boolean bool
type ProxyApi struct {
	Id     string
	New    bool
	Group  string // 组 （用于指定 host） \
	Path   string // 请求路径  		    |=> 用于唯一标识一个请求
	Method string // 请求方法 		   /
	//id 	   string
	Uri           string
	Cache         Boolean // 是否进行缓存
	CacheByURI    Boolean
	CacheWithBody Boolean
	ParentId      string
	//UnValid bool
	//Record bool
}

func (b Boolean) Value() bool {
	return bool(b)
}
func (b Boolean) toString() string {
	if b {
		return "true"
	} else {
		return "false"
	}
}

func (p ProxyApi) getId(all bool) string {
	if all {
		return fmt.Sprintf("#%s_$_*_$_%s#", p.Group, p.Method)
	} else {
		return fmt.Sprintf("#%s_$_%s_$_%s#", p.Group, p.Path, p.Method)
	}
}
func (p ProxyApi) ToString() string {
	return fmt.Sprintf(`Group = %s
Path = %s
Method = %s
Uri = %s
Cache = %s
CacheWithBody = %s
CacheByURI = %s
`, p.Group, p.Path, p.Method, p.Uri, p.Cache.toString(), p.CacheWithBody.toString(), p.CacheByURI.toString())
	//fmt.Println(fmt.Sprintf("Group = %s", p.Group))
	//fmt.Println(fmt.Sprintf("Path = %s", p.Path))
	//fmt.Println(fmt.Sprintf("Method = %s", p.Method))
	//fmt.Println(fmt.Sprintf("Uri = %s", p.Uri))
	//fmt.Println(fmt.Sprintf("Cache = %s", p.Cache.toString()))
}

func (ms *MainServer) AddProxyApi(api ProxyApi) {
	if ms.proxyApi == nil {
		ms.proxyApi = map[string]ProxyApi{}
	}
	//api.id = getId(api.Path,api.Method)
	//ms.ProxyApis = append(ms.ProxyApis, api)
	ms.proxyApi[api.getId(false)] = api
	tc.SaveProxyApi(api.getId(false), api)
}
func (ms *MainServer) getProxyApi(group, path, method string) (ProxyApi, bool) {
	p := ProxyApi{
		Group:  group,
		Path:   path,
		Method: method,
	}
	var check = func(all bool) bool {
		key := p.getId(all)
		if v, ok := ms.proxyApi[key]; ok {
			if v.getId(all) == key {
				p.CacheByURI = v.CacheByURI
				p.CacheWithBody = v.CacheWithBody
				p.Cache = v.Cache
				p.Uri = v.Uri
				return true
			}
		}
		return false
	}
	var ok bool
	if ok = check(false); ok {
		p.New = false
		return p, true
	}
	if ok = check(true); ok {
		p.New = true
		p.ParentId = p.getId(true)
		return p, true
	}
	return p, false
}

func (ms *MainServer) doProxy(w http.ResponseWriter, r *http.Request) {
	if ok := ms.SetCORS(&w, r); ok {
		(w).Write([]byte("ok"))
	} else {
		var returnOk = false
		p := strings.Split(r.URL.Path, "/")
		sind := 2
		if p[0] == "" {
			sind++
		}
		if len(p) > sind {
			p, ok := ms.getProxyApi(p[sind-1], strings.Join(p[sind:], "/"), r.Method)
			tc.SaveProxyApi(p.getId(false), p)
			p.ToString()
			if ok {
				returnOk = true
				s, cok := tc.Get(p.getId(false))
				if p.Cache.Value() && cok {
					fmt.Println("use cache")
					(w).Write([]byte(s))
				} else {
					ms.TAForwardHandler(w, r, p)
				}
			}
		}
		if !returnOk {
			w.Write([]byte(`{error: "no path"}`))
		}
	}
}

func (ms *MainServer) TAForwardHandler(writer http.ResponseWriter, request *http.Request, api ProxyApi) {
	u, err := url.Parse(api.Uri + api.Path)
	if nil != err {
		log.Println(err)
		return
	}

	proxy := httputil.ReverseProxy{
		Director: func(request *http.Request) {
			request.URL = u
		},
	}
	proxy.ModifyResponse = func(response *http.Response) error {
		//response.Header.Del("Access-Control-Allow-Origin")
		//buf := new(bytes.Buffer)
		////io.Copy(buf, response.Body)
		//buf.ReadFrom(response.Body)
		//respBytes := buf.String()
		//respString := string(respBytes)
		//
		////fmt.Println(respString)
		//
		//r := ioutil.NopCloser(strings.NewReader(respString))
		//response.Body = r
		//
		//return nil
		ms.ModifyResponse(response, api, request)
		return nil
	}

	proxy.ServeHTTP(writer, request)
}

func (ms MainServer) ModifyResponse(response *http.Response, api ProxyApi, request *http.Request) {
	response.Header.Del("Access-Control-Allow-Origin")
	buf := new(bytes.Buffer)
	//io.Copy(buf, response.Body)
	buf.ReadFrom(response.Body)
	respBytes := buf.String()
	respString := string(respBytes)

	if api.Cache {
		if api.CacheByURI {
			uris := strings.Split(request.RequestURI, "?")
			uri := ""
			if len(uris) == 2 {
				uri = uris[1]
			}
			if api.CacheWithBody {

			} else {
				tc.Save(api.getId(false)+uri, respString)
			}
		} else {
			tc.Save(api.getId(false), respString)
		}
	}

	//fmt.Println(respString)

	r := ioutil.NopCloser(strings.NewReader(respString))
	response.Body = r

}

func addProxyHandleFunc(ms *MainServer) {
	http.HandleFunc("/ProxyApi", func(writer http.ResponseWriter, request *http.Request) {
		ms.SetCORS(&writer, request)
		if request.Method != http.MethodPost {
			ReturnFail([]byte("Only post support")).Send(writer)
			return
		}
		var p Parmas
		p.init(request)
		if len(p.SourceParams) > 1 {
			if p.SourceParams[0] == GetProxyApiAll {
				ReturnSuccess(tc.GetAllProxyApi()).Send(writer)
			} else if p.SourceParams[0] == GetProxyApiOld {
				ReturnSuccess(tc.GetOldProxyApi()).Send(writer)
			} else if p.SourceParams[0] == GetProxyApiNew && len(p.SourceParams) > 2 {
				ReturnSuccess(tc.GetNewProxyApi(p.SourceParams[1])).Send(writer)
			} else {
				ReturnSuccess(map[string]ProxyApi{}).Send(writer)
			}
		} else {
			ReturnSuccess(map[string]ProxyApi{}).Send(writer)
		}
	})
}
