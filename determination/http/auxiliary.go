package http

import (
    "determination/determination/tool"
    "determination/determination/core"
    "determination/app/middleware"
    "reflect"
)
//阻塞用
func End(){
    select {};
}
//这里用这个init的问题是如果用init就不能用config了，加载的顺序问题,必须手动初始化
func HttpInit(){
	tool.LogInit()
    tool.DbInit()
}
//控制器调用
func controllerCall(controller string,method string,mh middleware.Http) bool{
    if tool.Config("Controller",tool.Capitalize(controller)+"Controller") != nil {
        rv := reflect.ValueOf(tool.Config("Controller",tool.Capitalize(controller)+"Controller")).Elem()
        //判断是否存在方法
        if(rv.MethodByName(tool.Capitalize(method)).IsValid() != false){
            //中间件核心代码
            if(middlewareCreate(func(request middleware.Http) interface{}{
                rv.FieldByName("W").Set(reflect.ValueOf(request.W))
                rv.FieldByName("R").Set(reflect.ValueOf(request.R))
                rv = rv.MethodByName(tool.Capitalize(method))
                return rv.Call([]reflect.Value{})
            },controller,method)(mh) != false){
                return true
            }
        }
    }
    return false
}
//中间件创建
func middlewareCreate(next middleware.Next,controller string,method string) middleware.Next{
    middlewareList := append(append(core.Middleware(controller,"@begin"),core.Middleware(controller,method)...),core.Middleware(controller,"@end")...)
    for i := len(middlewareList) - 1 ; i >= 0; i-- {
        next = middlewareMake(middlewareList[i],next)
    }
    return next
}
//中间件核心
func middlewareMake(thisfunc middleware.MakeMiddleware,next middleware.Next) middleware.Next{
    return func(request middleware.Http) interface{}{
        return thisfunc(request,next)
    }
}