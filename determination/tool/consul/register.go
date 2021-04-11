package consul

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"strconv"
)

func Register(ID string , Name string,Address string, Port string){
	intPort, _:= strconv.Atoi(Port)
    mjson,_ :=json.Marshal(map[string]interface{}{
	  "ID": ID,
	  "Name": Name,
	  "Address": Address,
	  "Port": intPort,          
	  "EnableTagOverride": false,
	  "Check":map[string]interface{}{             
	    "DeregisterCriticalServiceAfter": "90s",
	    "Http": "http://"+Address+":"+Port+"/", 
	    "Interval": "10s",
	  },
	})

	url := "http://106.55.38.162:8500/v1/agent/service/register"
	
	payload := strings.NewReader(string(mjson))
 
	req, _ := http.NewRequest("PUT", url, payload)
 
	req.Header.Add("Content-Type", "application/json")
 
	res, err := http.DefaultClient.Do(req)

	if err != nil{
		fmt.Println(Name+"自动注册失败错误原因",err)
		return
	}
	defer res.Body.Close()
	fmt.Println(Name+"自动注册成功")
}