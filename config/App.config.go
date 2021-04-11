package config

import "determination/determination/tool"

func (c Config) App() map[string]interface{}{
	return map[string]interface{}{
		
		//默认控制器
		"DEF_CONTROLLER":tool.Env("DEF_CONTROLLER","app"),
		//默认方法
		"DEF_METHOD":tool.Env("DEF_METHOD","index"),
		//日志最大缓存写入长度
		"LOG_IO_MSG_MAX_LENGTH":tool.Env("LOG_IO_MSG_MAX_LENGTH","10000"),
		//日志默认写入文件名
		"LOG_DEF_WRITE_FILENAME":tool.Env("LOG_DEF_WRITE_FILENAME","log"),
		//日志管道长度
		"LOG_CHAN_NUM":tool.Env("LOG_CHAN_NUM","10000"),

		//HTTP协议
		"HTTP_IP":tool.Env("HTTP_IP","0.0.0.0"),
		"HTTP_PORT":tool.Env("HTTP_PORT","9000"),
		//TCP协议
		"TCP_IP":tool.Env("TCP_IP","0.0.0.0"),
		"TCP_PORT":tool.Env("TCP_PORT","9001"),
		
		//consul注册发现
		"CONSUL_IP_PORT":tool.Env("CONSUL_IP_PORT","127.0.0.1:8500"),
		"CONSUL_HTTP_ID":tool.Env("CONSUL_HTTP_ID","determinationGoHttp"),
		"CONSUL_HTTP_NAME":tool.Env("CONSUL_HTTP_NAME","determinationGoHttp"),
		"CONSUL_TCP_ID":tool.Env("CONSUL_TCP_ID","determinationGoTcp"),
		"CONSUL_TCP_NAME":tool.Env("CONSUL_TCP_NAME","determinationGoTcp"),
		
	}
}