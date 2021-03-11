package middleware

import "net/http"

type MiddlewareRegister struct{}

type Next func(request interface{}) interface{}

type MakeMiddleware func(request interface{},next Next) interface{}

type Http struct{
	W http.ResponseWriter
	R *http.Request
}