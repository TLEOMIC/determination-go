package tool

import (
	"net/http"
	"encoding/json"
)

func EchoJson(w http.ResponseWriter,data interface{}){
    ret, _ := json.Marshal(&data)
    w.Header().Set("Content-Type", "text/html;charset=utf-8")
    w.Write(ret)
}