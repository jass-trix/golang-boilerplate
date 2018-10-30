package main

import (
	"github.com/go-ini/ini"
	"github.com/golang-boilerplate/pkg/server"
)

//Config is a struct for all config
type Config struct {
	Server server.Config
}

func init() {

}

//InitConfig is a function to instantiate configuration
func InitConfig() *Config {
	c := &Config{}
	err := ini.MapTo(c, "../configs/main.development.ini")
	if err != nil {
		return nil
	}
	return c
}
