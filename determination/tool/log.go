package tool

import (
	"os"
	"time"
	"io/ioutil"
)

var thisday string

var defLogName string

var ioMsgMaxLength int

var logChan map[string]chan string

var f map[string]*os.File
//这里不用init的问题是如果用init就不能用config了，加载的顺序问题,必须手动初始化
func Loginit(){
	fileInfoList,err := ioutil.ReadDir("./logs")
	if err != nil {
		panic("根目录找不到logs文件")
		return
	}
	logChan = make(map[string]chan string)
	f = make(map[string]*os.File)

	fileList := make([]string,len(fileInfoList),len(fileInfoList))
	for i := range fileInfoList {
		fileList[i] = string(fileInfoList[i].Name())
	}

	ioMsgMaxLength,_ = AppC("LOG_IO_MSG_MAX_LENGTH").(int)
	defLogName = AppC("LOG_DEF_WRITE_FILENAME").(string)
	go setGlobalVariable(fileList)

	num,_ := AppC("LOG_CHAN_NUM").(int)
	for i := range fileList{
		logChan[fileList[i]] = make(chan string,num)
		go logRun(logChan[fileList[i]],fileList[i])
	}
}

func setGlobalVariable(fileList []string){
	for{
		now := time.Now().String()[:10]
		if thisday == now{
			return
		}
		thisday = now
			for i := range fileList{
				f1, err := os.OpenFile("./logs/"+fileList[i]+"/"+thisday, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
				if err != nil {
			        logChan[defLogName] <- err.Error()+ "\r\n"
			    }else{
			    	f[fileList[i]] = f1
			    }
			}
	    SleepTillMidnight()
	}
}

//该方法在middleware/struct有引用方法
func Log(log string){
	L(defLogName,log)
}
//该方法在middleware/struct有引用方法
func L(LogName string,log string){
	logChan[LogName] <- log
}

func logRun(logC chan string,route string){
	var ioMsg string
	for{
		select {
			case msg := <-logC:
				ioMsg = msg + "\r\n"
				if len(ioMsg) > ioMsgMaxLength{
					logWrite(&ioMsg,route)
				}
			default:
				if ioMsg != "" {
					logWrite(&ioMsg,route)
				}else{
					ioMsg = <-logC + "\r\n"
				}
		}
	}
}
func logWrite(ioMsg *string,route string){
	f[route].Write([]byte(*ioMsg))
	*ioMsg = ""
}