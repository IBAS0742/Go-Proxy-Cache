package Server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type R struct {
	Error string
	Code  int
	Data  string
}

var RCodeSuccess = 200
var RCodeFail = 100
var RCodeError = 0

func ReturnSuccess(data interface{}) R {
	bts, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return R{
		Error: "",
		Code:  RCodeSuccess,
		Data:  string(bts),
	}
}

func ReturnFail(data interface{}) R {
	bts, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return R{
		Error: string(bts),
		Code:  RCodeFail,
		Data:  "",
	}
}

func (r R) Send(w http.ResponseWriter) {
	bts, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(bts))
}
