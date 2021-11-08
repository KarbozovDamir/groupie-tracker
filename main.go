package main

import (
	// "fmt"
	grtrack "grtrack/struct"
	"net/http"
	"time"
)

//include templates
func main() {
	http.HandleFunc("/main", grtrack.HundlerFirst)
	http.HandleFunc("/bands/", grtrack.HundlerSecond)
	http.ListenAndServe(":8080", nil)
}

func (s *Server) Run() error {
	handler := http.NewServeMux()

	srv := &http.Server{
		Addr:         s.Address,
		Handler:      handler,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
	}
	return srv.ListenAndServe()
}
