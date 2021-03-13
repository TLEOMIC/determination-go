package http 

import(
    "determination/determination/core"
    "determination/determination/tool"
    "determination/app/middleware"
	"net/http"
	"html"
	"reflect"
	"strings"
)
func WebRun(){
    tool.Loginit()
    tool.Dbinit()
    go goWebHttp(tool.AppC("PORT").(string))
}
func goWebHttp(port string){
    mux := http.NewServeMux()
    mux.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
        if r.URL.RequestURI() == "/favicon.ico" {
            return 
        }
        if !controllerCall(html.EscapeString(r.URL.Path)[1:],w,r) {
            w.WriteHeader(404)
        }
    })
    http.ListenAndServe(":"+port, mux)
}
func controllerCall(url string,w http.ResponseWriter,r *http.Request) bool{
	urlAnalysis := strings.Split(url, "/")
	if len(urlAnalysis) == 1{
        if urlAnalysis[0] == ""{
            urlAnalysis[0] = tool.AppC("DEF_CONTROLLER").(string)
        }
		urlAnalysis = append(urlAnalysis,tool.AppC("DEF_METHOD").(string))
	}else if urlAnalysis[1] == ""{
        urlAnalysis[1] = tool.AppC("DEF_METHOD").(string)
    }
    if tool.Config("Controller",tool.Capitalize(urlAnalysis[0])+"Controller") != nil {
        rv := reflect.ValueOf(tool.Config("Controller",tool.Capitalize(urlAnalysis[0])+"Controller")).Elem()
        //判断是否存在方法
        if(rv.MethodByName(tool.Capitalize(urlAnalysis[1])).IsValid() != false){
            //中间件核心代码
            next := func(request middleware.Http) interface{}{
                //最后再call
                rv.FieldByName("W").Set(reflect.ValueOf(request.W))
                rv.FieldByName("R").Set(reflect.ValueOf(request.R))
                rv = rv.MethodByName(tool.Capitalize(urlAnalysis[1]))
                return rv.Call([]reflect.Value{})
            }
            middlewareList := core.Middleware(urlAnalysis[0],urlAnalysis[1])
            for i := len(middlewareList) - 1 ; i >= 0; i-- {
                next = middlewareMake(middlewareList[i],next)
            }
            next(middleware.Http{W:w,R:r})
            return true
        }
    }
    return false
}

func middlewareMake(thisfunc middleware.MakeMiddleware,next middleware.Next) middleware.Next{
    return func(request middleware.Http) interface{}{
        return thisfunc(request,next)
    }
}