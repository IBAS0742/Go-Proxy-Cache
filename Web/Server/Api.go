package Server

import "net/http"

type Api struct {
	MethodName string
	Dear       func(w http.ResponseWriter, params Parmas, r *http.Request)
}
