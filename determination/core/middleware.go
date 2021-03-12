package core

import (
	"determination/app/middleware"
	"io/ioutil"
	"reflect"
	"strings"
	"determination/determination/tool"
)

var middlewareMap map[string]interface{}
var middlewareRe reflect.Value

func init(){
	middlewareRe = reflect.ValueOf(new(middleware.MiddlewareRegister)).Elem()
	middlewareMap = make(map[string]interface{})
	makeMiddleware()
}
func Middleware(key1 string,key2 string) []middleware.MakeMiddleware{
	return middlewareMap[tool.Capitalize(key1)].(map[string][]middleware.MakeMiddleware)[tool.Capitalize(key2)]
}
func middlewareCall(methods string) map[string][]middleware.MakeMiddleware{
	return middlewareRe.MethodByName(methods).Call([]reflect.Value{})[0].Interface().(map[string][]middleware.MakeMiddleware)
}
func makeMiddleware(){
	fileInfoList,err := ioutil.ReadDir("./app/middleware")
	if err != nil {
		panic("app目录找不到middleware文件")
		return
	}
	var Split []string
	for i := range fileInfoList {
		Split = strings.Split(string(fileInfoList[i].Name()), ".")
		if Split[1] == "middleware" {
			middlewareMap[tool.Capitalize(Split[0])] = middlewareCall(tool.Capitalize(Split[0])+"MiddlewareRegister")
		}
	}
}