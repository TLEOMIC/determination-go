package http 

import(
    "determination/determination/tool"
    "determination/app/middleware"
    "net/http"
	"html"
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
        controller,method :=makeControllerAndMethod(html.EscapeString(r.URL.Path)[1:],"/")
        if controllerCall(controller,method,middleware.Http{W:w,R:r}) == false {
            w.WriteHeader(404)
        }
    })
    http.ListenAndServe(":"+port, mux)
}
