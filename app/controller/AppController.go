package controller

import (
	"fmt"
	
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
	c.EchoJson(map[string]string{"code":"0","msg":"这是首页"})
}