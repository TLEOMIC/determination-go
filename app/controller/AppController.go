package controller

import (
	"fmt"
	"determination/determination/tool"
)

type AppController struct{
	Controller
}

func (c AppController) Test2(){
	fmt.Println(c)
}
func (c AppController) Test(){
	// fmt.Println(c)
	fmt.Println(tool.AppC("LOG_IO_MSG_MAX_LENGTH"))
}
func (c AppController) Index(){
	c.EchoJson(map[string]string{"code":"0","msg":"这是首页"})
}