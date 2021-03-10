目录结构
app 	主目录
	|_ controller 控制器目录
config  配置目录
determination 框架核心目录
-----------------------------------------------------
/app/controller说明

结构体必须继承 Controller
结构体名和文件名可以不一致
结构体必须在/config内的Controller.config.go中声明
-----------------------------------------------------
/config说明
import "determination/determination/config"
会自动读取所有在/config的包含 \*.config.go的文件

返回app配置内指定key的数据
AppC(key string) interface{}

返回指定key1文件配置内key2的数据
Config(key1 string,key2 string) interface{}
-----------------------------------------------------
api说明

import "determination/determination/tool"

首字母大写
Capitalize(str string) string 

页面输出json
EchoJson(w http.ResponseWriter,data interface{})

判断是否为空
Empty(params interface{}) bool

读取env文件，key是键 value是如果env内不存在则使用的默认值
Env(key string,value string) string