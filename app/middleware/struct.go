package middleware

import "net/http"

type MiddlewareRegister struct{}

type Next func(request Http) interface{}

type MakeMiddleware func(request Http,next Next) interface{}

type Http struct{
	W http.ResponseWriter
	R *http.Request
}