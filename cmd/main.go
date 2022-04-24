package main

import (
	"fmt"
	"hemli/internal/config"
	"hemli/internal/server"
	"net/http"
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
	s.MountHandlers()

	addr := fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	fmt.Printf("Application started on: http://%s\n", addr)
	http.ListenAndServe(addr, s.Router)
}
