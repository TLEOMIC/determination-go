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
	fmt.Println(c)
}
func (c AppController) Index(){
	tool.EchoJson(c.W,map[string]string{"code":"0","msg":"这是首页"})
}