package config

import "determination/determination/tool"

func (c Config) App() map[string]interface{}{
	return map[string]interface{}{
		"IP":tool.Env("IP","127.0.0.1"),
		"PORT":tool.Env("PORT","9000"),
	}
}