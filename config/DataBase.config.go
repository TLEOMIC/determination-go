package config

import "determination/determination/tool"
//框架目前只实现mysql
func (c Config) DataBase() map[string]interface{}{
	return map[string]interface{}{
		"mycat":map[string]string{
			"IP":tool.Env("DB_HOST","127.0.0.1"),
			"PORT":tool.Env("DB_PORT","3306"),
			"DATABASE":tool.Env("DB_DATABASE","mycat"),
			"USER":tool.Env("DB_USER","root"),
			"PASSWORD":tool.Env("DB_PASSWORD","root"),
			"NETWORK_PROTOCOL":tool.Env("DB_NETWORK_PROTOCOL","tcp"),
			"CONN_MAX_LIFE_TIME":tool.Env("DB_CONN_MAX_LIFE_TIME","3"),
			"MAX_IDLE_CONNS":tool.Env("DB_MAX_IDLE_CONNS","5"),
			"MAX_OPEN_CONNS":tool.Env("DB_MAX_OPEN_CONNS","10"),
		},
		"db2":map[string]string{
			"IP":tool.Env("DB_HOST","127.0.0.1"),
			"PORT":tool.Env("DB_PORT","3306"),
			"DATABASE":tool.Env("DB_DATABASE","db2"),
			"USER":tool.Env("DB_USER","root"),
			"PASSWORD":tool.Env("DB_PASSWORD","root"),
			"NETWORK_PROTOCOL":tool.Env("DB_NETWORK_PROTOCOL","tcp"),
			"CONN_MAX_LIFE_TIME":tool.Env("DB_CONN_MAX_LIFE_TIME","3"),
			"MAX_IDLE_CONNS":tool.Env("DB_MAX_IDLE_CONNS","5"),
			"MAX_OPEN_CONNS":tool.Env("DB_MAX_OPEN_CONNS","10"),
		},
	}
}