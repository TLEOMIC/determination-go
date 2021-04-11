package consul

import (
	"determination/determination/tool"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"strconv"
)

func Register(ID string , Name string,Address string, Port string,Type string){
	intPort, _:= strconv.Atoi(Port)
	consulJson := map[string]interface{}{
	  "ID": ID,
	  "Name": Name,
	  "Address": Address,
	  "Port": intPort,          
	  "EnableTagOverride": false,
	}
	
	switch Type {
		case "tcp" :
			consulJson["Check"] = map[string]interface{}{             
			    "DeregisterCriticalServiceAfter": "90s",
			    "tcp": Address+":"+Port, 
			    "Interval": "10s",
			}
		case "http" :
			consulJson["Check"] = map[string]interface{}{             
			    "DeregisterCriticalServiceAfter": "90s",
				"http": "http://"+Address+":"+Port+"/", 
				"Interval": "10s",
			}
	}
    mjson,_ :=json.Marshal(consulJson)

	url := "http://"+tool.AppC("CONSUL_IP_PORT").(string)+"/v1/agent/service/register"
	
	payload := strings.NewReader(string(mjson))
 
	req, _ := http.NewRequest("PUT", url, payload)
 
	req.Header.Add("Content-Type", "application/json")
 
	res, err := http.DefaultClient.Do(req)

	if err != nil{
		fmt.Println(Name+"自动注册失败错误原因",err)
		return
	}
	defer res.Body.Close()
	fmt.Println(Name+"自动注册成功,请先检查一下consul是否有该服务")
}