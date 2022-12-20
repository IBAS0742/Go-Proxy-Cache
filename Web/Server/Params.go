package Server

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Parmas struct {
	MethodName   string
	Params       map[string]string
	SourceParams []string
}

func (p *Parmas) init(r *http.Request) {
	var err error
	var errMsg string
	if r.Method == Methods["GET"] {
		sp := strings.Split(r.RequestURI, "?")
		if len(sp) == 2 {
			pars := strings.Split(sp[1], "&")
			for _, _p := range pars {
				sp = strings.Split(_p, "=")
				if len(sp) == 2 {
					sp[0] = strings.ToLower(sp[0])
					if sp[0] == "methodname" {
						p.MethodName = sp[1]
					} else {
						p.Params[sp[0]] = sp[1]
					}
					if sp[0] == "sourceparams" {
						p.SourceParams = strings.Split(sp[1], ",")
					}
				}
			}
		} else {
			errMsg = "Method = [GET] and can not get query string"
		}
	} else if r.Method == Methods["POST"] {
		err = json.NewDecoder(r.Body).Decode(&p)
	} else {
		errMsg = "Method = [" + r.Method + "] can not be found"
	}
	if err != nil {
		panic(err)
	} else if errMsg != "" {
		panic(errMsg)
	}
}

type Mss map[string]string

func (mss Mss) GetByKey(key string) (string, bool) {
	if value, ok := mss[key]; ok {
		return value, ok
	} else {
		return "", false
	}
}
