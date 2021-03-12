package core

import (
	"determination/config"
	"io/ioutil"
	"reflect"
	"strings"
	"determination/determination/tool"
)
func init(){
	configMap := make(map[string]interface{})
	configRe := reflect.ValueOf(new(config.Config)).Elem()

	fileInfoList,err := ioutil.ReadDir("./config")
	if err != nil {
		panic("根目录找不到config文件")
		return
	}
	var Split []string
	for i := range fileInfoList {
		Split = strings.Split(string(fileInfoList[i].Name()), ".")
		if Split[1] == "config" {
			configMap[tool.Capitalize(Split[0])] = configRe.MethodByName(tool.Capitalize(Split[0])).Call([]reflect.Value{})[0].Interface().(map[string]interface{})
		}
	}
	tool.SetConfigMap(configMap)
}