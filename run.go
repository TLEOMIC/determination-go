package main

import "determination/determination/http"

func main(){
	//加载框架核心
	http.HttpInit()
	//开启web服务
	http.WebRun()
	//阻塞
	http.End()
}
