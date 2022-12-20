package MDB

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"unsafe"
)

type ApiConfig struct {
	Group         string
	Uri           string
	Method        string // http.MethodGet,
	Header        map[string]string
	Params        map[string]string
	Cache         bool
	ReturnType    string
	DefaultReturn string
	DelayTime     int
}

const (
	ReturnTypeString = "string"
	ReturnTypeJson   = "json"
	ReturnTypeFile   = "file"
)

var emptyBody = strings.NewReader("{}")

func (api ApiConfig) buildBody() io.Reader {
	if api.Params == nil {
		return emptyBody
	}
	bodyJson, err := json.Marshal(api.Params)
	if err != nil {
		panic(err)
		return emptyBody
	}
	return strings.NewReader(string(bodyJson))
}
func (api ApiConfig) setHeader(r *http.Request) {
	for k, v := range api.Header {
		r.Header.Set(k, v)
	}
}
func (api ApiConfig) Do() string {
	req, err := http.NewRequest(
		api.Method,
		api.Uri,
		api.buildBody(),
	)
	if err != nil {
		panic(err)
	}
	api.setHeader(req)
	resp, err := http.DefaultClient.Do(req)

	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	//byte数组直接转成string,优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	//fmt.Println()
	return *str
}
