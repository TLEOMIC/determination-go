package middleware

import "fmt"

func (Mr MiddlewareRegister) AppMiddlewareRegister() map[string][]MakeMiddleware{
	return map[string][]MakeMiddleware{
		"Test":{midd1,midd2},
	}
}
func midd1(request Http,next Next) interface{}{
	fmt.Println("run2")
	fmt.Println(request.W)
	return next(request)
}
func midd2(request Http,next Next) interface{}{
	fmt.Println("run3")
	newrequest := next(request)
	fmt.Println("run4")
	return newrequest
}