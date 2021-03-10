package config

import (
	"determination/config"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"determination/determination/tool"
)

var configMap map[string]string
var configRe reflect.Value

func init(){
	configRe = reflect.ValueOf(new(config.Config)).Elem()
	configMap = make(map[string]string)
	makeConfig()
}
func Config(){
	fmt.Println(configMap)
}

func configCall(methods string) map[string]string{
	return configRe.MethodByName(methods).Call([]reflect.Value{})[0].Interface().(map[string]string)
}

func makeConfig(){

	// config.Call()
	// whattype := configCall("App")
	// fmt.Println(whattype["IP"])

	// return 
	fileInfoList,err := ioutil.ReadDir("./config")
	if err != nil {
		fmt.Println("err")
		return
	}
	var Split []string
	for i := range fileInfoList {
		Split = strings.Split(string(fileInfoList[i].Name()), ".")
		if Split[1] == "config" {
			configMap = tool.Collection(configMap,configCall(Split[0]))
		}
	}


	
}