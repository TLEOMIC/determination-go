package controller

import (
	"net/http"
	"determination/determination/tool"
)

type Controller struct{
	W http.ResponseWriter
	R *http.Request
	Tcp interface{}
}
func (c Controller) EchoJson(m interface{}){
	tool.EchoJson(c.W,m)
}
