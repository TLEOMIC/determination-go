package http 

import(
    "determination/determination/tool"
	"determination/determination/config"
	"net/http"
	"html"
	"reflect"
	"strings"
)

func Run(){
    go goHttp(config.AppC("PORT").(string));
    select {};
}
func goHttp(port string){
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
            urlAnalysis[0] = tool.Env("DEF_CONTROLLER","index")
        }
		urlAnalysis = append(urlAnalysis,tool.Env("DEF_METHOD","index"))
	}else if urlAnalysis[1] == ""{
        urlAnalysis[1] = tool.Env("DEF_METHOD","index")
    }

	rv := reflect.ValueOf(config.Config("Controller",tool.Capitalize(urlAnalysis[0])+"Controller")).Elem()

    rv.FieldByName("W").Set(reflect.ValueOf(w))
    rv.FieldByName("R").Set(reflect.ValueOf(r))

    rv = rv.MethodByName(tool.Capitalize(urlAnalysis[1]))

    if(rv.IsValid() != false){
    	rv.Call([]reflect.Value{})
        return true
    }else{
    	return false
    }
}