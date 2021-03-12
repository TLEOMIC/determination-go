package tool

import (
	"os"
	"time"
	"strconv"
)

var thisday string

var ioMsgMaxLength int

var logChan chan string

var f *os.File

func init(){
	num,_ := strconv.Atoi(Env("LOG_CHAN_NUM","10000"))
	ioMsgMaxLength,_ = strconv.Atoi(Env("LOG_IO_MSG_MAX_LENGTH","10000"))
	logChan = make(chan string,num)
	go setGlobalVariable()
	go logRun(logChan)
}

func setGlobalVariable(){
	for{
		now := time.Now().String()[:10]
		if thisday == now{
			return
		}
		thisday = now
		f1, err := os.OpenFile("./log/"+thisday, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
		if err != nil {
	        logChan <- err.Error()+ "\r\n"
	    }else{
	    	f = f1
	    }
	    year, month, day := time.Now().Add(24*time.Hour).Date()
	    time.Sleep(time.Duration(time.Date(year, month, day, 0, 0, 0, 0, time.Local).Unix()-time.Now().Unix())*time.Second)
	}
	
}

//该方法在middleware/struct有引用方法
func Log(log string){
	logChan <- log
}
func logRun(logC chan string){
	var ioMsg string
	for{
		select {
			case msg := <-logC:
				ioMsg = msg + "\r\n"
				if len(ioMsg) > ioMsgMaxLength{
					logWrite(&ioMsg)
				}
			default:
				if ioMsg != "" {
					logWrite(&ioMsg)
				}
		}
	}
}
func logWrite(ioMsg *string){
	f.Write([]byte(*ioMsg))
	*ioMsg = ""
}