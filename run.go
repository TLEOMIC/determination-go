package main

import "determination/determination/http"

func main(){
	//加载框架核心
	http.HttpInit()
	//开启web服务
	http.WebRun()
	//开启tcp服务
	http.TcpRun()
	//阻塞
	http.End()
}
