package config

import (
	"determination/config"
	"io/ioutil"
	"reflect"
	"strings"
	"determination/determination/tool"
)

var configMap map[string]interface{}
var configRe reflect.Value

func init(){
	configRe = reflect.ValueOf(new(config.Config)).Elem()
	configMap = make(map[string]interface{})
	makeConfig()
}
func AppC(key string) interface{}{
	return Config("App",key)
}
func Config(key1 string,key2 string) interface{}{
	return configMap[tool.Capitalize(key1)].(map[string]interface{})[key2]
}
func configCall(methods string) map[string]interface{}{
	return configRe.MethodByName(methods).Call([]reflect.Value{})[0].Interface().(map[string]interface{})
}
func makeConfig(){
	fileInfoList,err := ioutil.ReadDir("./config")
	if err != nil {
		panic("根目录找不到config文件")
		return
	}
	var Split []string
	for i := range fileInfoList {
		Split = strings.Split(string(fileInfoList[i].Name()), ".")
		if Split[1] == "config" {
			configMap[tool.Capitalize(Split[0])] = configCall(tool.Capitalize(Split[0]))
		}
	}
}