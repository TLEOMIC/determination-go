package controller

import (
	"fmt"
	"determination/determination/tool"
)

type AppController struct{
	Controller
}

func (c AppController) Test2(){

		rows, err := tool.Db("mycat").Query("select a from db1")
		if err != nil {  
			fmt.Println(err)
		}
		var a string
		for rows.Next(){
			err := rows.Scan(&a)
			if err != nil {  
				fmt.Println(err)
			}
			fmt.Println(a)
		}
		rows.Close()

}
func (c AppController) Test(){
	// fmt.Println(c)
	fmt.Println(tool.AppC("LOG_IO_MSG_MAX_LENGTH"))
}
func (c AppController) Index(){
	c.EchoJson(map[string]string{"code":"0","msg":"这是首页"})
}