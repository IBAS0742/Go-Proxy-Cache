package main

import (
	"bytes"
	"fmt"
	"go-proxy-cache/Web/Server"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	server := Server.MainServer{
		Port: 8090,
		CORS: Server.DefaultCors,
		//ImportStatic: nil,
		//StaticPath:   nil,
		ImportStatic: func() {},
	}
	server.AddProxyApi(Server.ProxyApi{
		Group: "table",
		//Path:       "system/user/list",
		Path:       "*",
		Method:     http.MethodGet,
		Uri:        "http://172.20.109.155:8080/",
		Cache:      true,
		CacheByURI: true,
	})
	server.AddProxyApi(Server.ProxyApi{
		Group: "table",
		//Path:       "system/user/list",
		Path:       "captchaImage",
		Method:     http.MethodGet,
		Uri:        "http://172.20.109.155:8080/",
		Cache:      false,
		CacheByURI: false,
	})
	server.AddProxyApi(Server.ProxyApi{
		Group: "table",
		//Path:       "system/user/list",
		Path:       "login",
		Method:     http.MethodPost,
		Uri:        "http://172.20.109.155:8080/",
		Cache:      false,
		CacheByURI: false,
	})
	server.RunServer()
}

func main1() {
	server := Server.MainServer{
		Port: 8090,
		CORS: Server.DefaultCors,
		//ImportStatic: nil,
		//StaticPath:   nil,
		ImportStatic: func() {},
		Proxy: func(writer http.ResponseWriter, request *http.Request) {
			p := strings.Split(request.URL.Path, "/")
			sind := 2
			if p[0] == "" {
				sind++
			}
			if len(p) > sind {
				if p[sind-1] == "table" {
					TAForwardHandler(writer, request, strings.Join(p[sind:], "/"))
					return
				}
			}
			writer.Write([]byte(`{error: "no path"}`))
			return
		},
	}
	server.RunServer()
}

func TAForwardHandler(writer http.ResponseWriter, request *http.Request, urlPath string) {
	u, err := url.Parse("http://172.20.109.155:8080/" + urlPath)
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
		response.Header.Del("Access-Control-Allow-Origin")
		buf := new(bytes.Buffer)
		//io.Copy(buf, response.Body)
		buf.ReadFrom(response.Body)
		respBytes := buf.String()
		respString := string(respBytes)
		fmt.Println(respString)

		r := ioutil.NopCloser(strings.NewReader(respString))
		response.Body = r

		return nil
	}

	proxy.ServeHTTP(writer, request)
}
