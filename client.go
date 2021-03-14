package main

import (
   "fmt"
   "net/rpc/jsonrpc"
)

type Goods struct {
   Id   int
   Name string
}

type Params struct {
   Id   int
   Name string
}

func main() {
   
   conn, _ := jsonrpc.Dial("tcp", "127.0.0.1:9001")
   defer conn.Close()
   var data string
   var p Params
   p.Id = 1
   // 等价于这行{"method":"tcp.Test","params":[{"Id":0,"Name":""}],"id":0}
   err := conn.Call("tcp.Test", p, &data)
   if err != nil {
      fmt.Println("err : ", err)
   }
   fmt.Println("data : ", data)
}
