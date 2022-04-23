package main

import (
	"fmt"
	"hemli/internal/config"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	conf = config.NewConfig()
)

func init() {
	// Load environment variables and configuration parameters
	conf.Init()
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello from go and chi"))
	})

	addr := fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	fmt.Printf("Application started on: http://%s\n", addr)
	http.ListenAndServe(addr, r)
}
