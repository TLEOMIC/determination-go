package middleware

import (
	"net/http"
	"determination/determination/tool"
)


type MiddlewareRegister struct{}

type Next func(request Http) interface{}

type MakeMiddleware func(request Http,next Next) interface{}

type Http struct{
	W http.ResponseWriter
	R *http.Request
}

func Log(log string){
	tool.Log(log)
}
func L(LogName string,log string){
	tool.L(LogName,log)
}