运行
根目录运行 go run run.go即可
端口在env内
本框架无路由按着控制器+方法进入
如访问127.0.0.1:9000/app/test则访问的是AppController结构体的test方法
127.0.0.1:9000默认进入的方法在env内配置
-----------------------------------------------------
目录结构
app 	主目录
	|_ controller 控制器目录
	|_ middleware 中间件目录
config  配置目录
determination 框架核心目录
-----------------------------------------------------
/app/controller说明

结构体必须继承 Controller
结构体名和文件名可以不一致
结构体必须在/config内的Controller.config.go中声明
-----------------------------------------------------
/app/middleware说明

框架运行会自动注册/app/middleware目录下所有\*.middleware.go的文件
方法名称前缀与文件名必须一致,如要加一个Test的配置 方法后缀为MiddlewareRegister
func (Mr MiddlewareRegister) TestMiddlewareRegister() map[string][]MakeMiddleware{
	return map[string][]MakeMiddleware{
		"Test":{demo1,demo2},
	}
}
文件名则是Test.middleware.go

中间件使用格式
next的位置会影响是在方法之前执行还是方法之后执行，具体使用和laravel框架的中间件几乎一致
func demo1(request interface{},next Next) interface{}{
	return next(request)
}
request使用时需要转换成Http才能使用（这个之后应该会修复
-----------------------------------------------------
/config说明
import "determination/determination/config"
会自动读取所有在/config的包含 \*.config.go的文件
方法与文件名必须一致,如要加一个test的配置
func (c Config) Test() map[string]interface{}{
	return map[string]interface{}{}
}
文件名则是Test.config.go

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