package controller

import "net/http"

type Controller struct{
	W http.ResponseWriter
	R *http.Request
}
