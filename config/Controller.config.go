package config

import (
	"determination/app/controller"
)

func (c Config) Controller() map[string]interface{}{
	return map[string]interface{}{
		"AppController":new(controller.AppController),
	}
}