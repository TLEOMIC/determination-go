package http 

import(
    "determination/determination/tool"
    "determination/app/middleware"
    "net/http"
	"html"
	"strings"
)
func WebRun(){

    go goWebHttp(tool.AppC("PORT").(string))
}
func goWebHttp(port string){
    mux := http.NewServeMux()
    mux.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
        if r.URL.RequestURI() == "/favicon.ico" {
            return 
        }       
        controller,method :=webHttpMakeControllerAndMethod(html.EscapeString(r.URL.Path)[1:])
        if !controllerCall(controller,method,middleware.Http{W:w,R:r}) {
            w.WriteHeader(404)
        }
    })
    http.ListenAndServe(":"+port, mux)
}
func webHttpMakeControllerAndMethod(url string) (string,string){
    urlAnalysis := strings.Split(url, "/")
    controller := urlAnalysis[0]
    method := urlAnalysis[1]
    if len(urlAnalysis) == 1{
        if controller == ""{
            controller = tool.AppC("DEF_CONTROLLER").(string)
        }
     urlAnalysis = append(urlAnalysis,tool.AppC("DEF_METHOD").(string))
    }else if method == ""{
        method = tool.AppC("DEF_METHOD").(string)
    }
    return controller,method
}
