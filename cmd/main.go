package main

import (
	"hemli/internal/config"
	"hemli/internal/server"
)

var (
	conf = config.NewConfig()
)

func init() {
	// Load environment variables and configuration parameters
	conf.Init()
}

func main() {
	s := server.CreateNewServer()
	s.Run(conf.Host, conf.Port)
}
