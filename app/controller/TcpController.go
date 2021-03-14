package controller

import (
	"fmt"
)

type TcpController struct{
	Controller
}

func (c TcpController) Test() interface{}{
	fmt.Println("join")
	return "hello c"
}
