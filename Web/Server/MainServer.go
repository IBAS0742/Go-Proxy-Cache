package Server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var Methods = map[string]string{
	"POST":    "POST",
	"GET":     "GET",
	"OPTIONS": "OPTIONS",
	"PUT":     "PUT",
	"DELETE":  "DELETE",
}

type Cors struct {
	Origin  string
	Methods []string
}

var DefaultCors = Cors{
	Origin: "*",
	Methods: []string{
		Methods["POST"],
		Methods["GET"],
		Methods["OPTIONS"],
		Methods["PUT"],
		Methods["DELETE"],
	},
}

type MainServer struct {
	run  bool
	Port int
	CORS Cors
	apis []Api

	ImportStatic func()

	// map[string] string {
	//    "view": "C:\\html\\"
	// }
	StaticPath map[string]string
	//ProxyApis  []ProxyApi
	proxyApi map[string]ProxyApi
	Proxy    func(w http.ResponseWriter, r *http.Request)
}

var tc = TmpCache{}

func (ms *MainServer) RunServer() {
	if ms.run {
		return
	}
	ms.run = true
	srv := &http.Server{
		Addr:           ":" + strconv.Itoa(ms.Port),
		Handler:        nil,
		ReadTimeout:    time.Duration(5) * time.Minute,
		WriteTimeout:   time.Duration(5) * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}
	var err error
	ms.ImportStatic()

	for k, v := range ms.StaticPath {
		p := k
		if k[0] != '/' {
			p = "/" + k
		}
		if k[len(k)-1] == '/' {
			p += "/"
		}
		http.Handle(k, http.StripPrefix(k, http.FileServer(http.Dir(v))))
	}

	http.HandleFunc("/proxy/", ms.doProxy)
	//http.HandleFunc("/proxy/", func(w http.ResponseWriter, r *http.Request) {
	//	if ok := ms.SetCORS(&w, r); ok {
	//		w.Write([]byte("ok"))
	//	} else {
	//		ms.Proxy(w, r)
	//	}
	//})
	http.HandleFunc("/api", ms.action())
	addProxyHandleFunc(ms)
	//statikFS, err := fs.New()
	//http.Handle("/leaflet171/", http.StripPrefix("/leaflet171/", http.FileServer(statikFS)))
	fmt.Println("【浏览/view】 http://localhost:" + strconv.Itoa(ms.Port))
	err = srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}

}

func (ms *MainServer) SetCORS(w *http.ResponseWriter, r *http.Request) bool {
	setupCORS := func(w *http.ResponseWriter) {
		//(*w).Header().Set("Access-Control-Allow-Origin", ms.CORS.Origin)
		(*w).Header().Set("Access-Control-Allow-Methods", strings.Join(ms.CORS.Methods, ","))
		(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		(*w).Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
		//(*w).Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		(*w).Header().Set("content-type", "application/json") //返回数据格式是json
	}
	setOptionReturn := func(w http.ResponseWriter, r *http.Request) bool {
		return false
	}
	for _, m := range ms.CORS.Methods {
		if m == Methods["OPTIONS"] {
			setOptionReturn = func(w http.ResponseWriter, r *http.Request) bool {
				if r.Method == Methods["OPTIONS"] {
					//w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
					return true
				} else {
					return false
				}
			}
		}
	}
	setupCORS(w)
	return setOptionReturn(*w, r)
}

func (ms *MainServer) AddApi(api Api) {
	ms.apis = append(ms.apis, api)
}

func (ms *MainServer) AddApis(apis []Api) {
	for _, api := range apis {
		ms.apis = append(ms.apis, api)
	}
}

func (ms *MainServer) action() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if ok := ms.SetCORS(&w, r); ok {
			return
		}
		var p Parmas
		p.init(r)
		ms.doApi(p, w, r)
	}
}

func (ms *MainServer) doApi(p Parmas, w http.ResponseWriter, r *http.Request) {
	for _, api := range ms.apis {
		if api.MethodName == p.MethodName {
			api.Dear(w, p, r)
			return
		}
	}
	fmt.Fprintf(w, "api[MethodName]:{"+p.MethodName+"} can not be found")
}
