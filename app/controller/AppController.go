package controller

import "fmt"

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
	fmt.Println(c)
}