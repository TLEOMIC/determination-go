package config

import "determination/determination/tool"

func (c Config) App() map[string]interface{}{
	return map[string]interface{}{
		"IP":tool.Env("IP","127.0.0.1"),
		"PORT":tool.Env("PORT","9000"),
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

	}
}