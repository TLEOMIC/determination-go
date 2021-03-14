package http

import (
    "determination/determination/tool"
    "determination/app/middleware"
    "fmt"
    "net"
    "bufio"
    "io"
    "encoding/json"
)

func TcpRun(){
    go goTcp(tool.AppC("TCP_PORT").(string))
}
func goTcp(port string){
    ln, err := net.Listen("tcp", ":"+port)
    if err != nil {
        panic("tcp启动失败:"+err.Error())
    }
    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Listen.Accept failed,err:",err)
            continue
        }
        go tcpHandle(conn)
    }
}
func tcpHandle(conn net.Conn){
    defer conn.Close() 
    reader   := bufio.NewReader(conn)
    data,err := reader.ReadSlice('\n')
    if err != nil {
        if err != io.EOF {
            fmt.Println(err)
        }
    }
    //consul的心跳处理处理
    if(string(data)==""){
        return
    }
    var m map[string]interface{}
    err = json.Unmarshal([]byte(data), &m)
    if err != nil {
        fmt.Println(err)
    }
    var Rdata interface{}
    if m["params"] != nil{
        controller,method :=makeControllerAndMethod(m["method"].(string),".")
        Rdata = controllerCall(controller,method,middleware.Http{Tcp:m["params"]})
    }
    if Rdata == false {
        Rdata = "404"
    }
    mjson,_ :=json.Marshal(map[string]interface{}{"id":0,"result":Rdata,"err":nil})
    _, err = conn.Write([]byte(string(mjson)+"\n"))
    if err != nil {
        fmt.Println(err)
    }
}