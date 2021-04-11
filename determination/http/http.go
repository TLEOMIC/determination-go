package http 

import(
    "determination/determination/tool"
    "determination/determination/tool/consul"
    "determination/app/middleware"
    "net/http"
	"html"
)
func WebRun(){
    go goWebHttp(tool.AppC("HTTP_IP").(string),tool.AppC("HTTP_PORT").(string))
    go consul.Register(tool.AppC("CONSUL_HTTP_ID").(string),tool.AppC("CONSUL_HTTP_NAME").(string),tool.GetMyIp(),tool.AppC("HTTP_PORT").(string),"http")
}
func goWebHttp(ip string,port string){
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
    http.ListenAndServe(ip+":"+port, mux)
}
