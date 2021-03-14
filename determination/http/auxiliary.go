package http

import (
    "determination/determination/tool"
    "determination/determination/core"
    "determination/app/middleware"
    "reflect"
    "strings"
)
//阻塞用
func End(){
    select {};
}
//这里用这个init的问题是如果用init就不能用config了，加载的顺序问题,必须手动初始化
func HttpInit(){
    if tool.Env("DB_INIT","true") != "false"{
        tool.DbInit()
    }
	tool.LogInit()
}
//控制器调用
func controllerCall(controller string,method string,mh middleware.Http) interface{}{
    if tool.Config("Controller",tool.Capitalize(controller)+"Controller") != nil {
        rv := reflect.ValueOf(tool.Config("Controller",tool.Capitalize(controller)+"Controller")).Elem()
        //判断是否存在方法
        if(rv.MethodByName(tool.Capitalize(method)).IsValid() != false){
            //中间件核心代码
            Rdata := middlewareCreate(func(request middleware.Http) interface{}{
                if request.W != nil {
                    rv.FieldByName("W").Set(reflect.ValueOf(request.W))
                }
                if request.R != nil {
                    rv.FieldByName("R").Set(reflect.ValueOf(request.R))
                }
                if request.Tcp != nil {
                    rv.FieldByName("Tcp").Set(reflect.ValueOf(request.Tcp))
                }
                rv = rv.MethodByName(tool.Capitalize(method))
                rvData := rv.Call([]reflect.Value{})
                if len(rvData) == 1{
                    return rvData[0].Interface()
                }
                return rvData 
            },controller,method)(mh)
            return Rdata
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
func makeControllerAndMethod(url string,separator string) (string,string){
    if url == ""{
        return tool.AppC("DEF_CONTROLLER").(string),tool.AppC("DEF_METHOD").(string)
    }
    urlAnalysis := strings.Split(url, separator)
    controller := urlAnalysis[0]
    method := urlAnalysis[1]
    if len(urlAnalysis) == 1{
     urlAnalysis = append(urlAnalysis,tool.AppC("DEF_METHOD").(string))
    }else if method == ""{
        method = tool.AppC("DEF_METHOD").(string)
    }
    return controller,method
}
