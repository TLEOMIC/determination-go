package tool

var configMap map[string]interface{}

func SetConfigMap(newconfigMap map[string]interface{}){
	configMap = newconfigMap
}
func AppC(key string) interface{}{
	return Config("App",key)
}
func Config(key1 string,key2 string) interface{}{
	return configMap[Capitalize(key1)].(map[string]interface{})[key2]
}