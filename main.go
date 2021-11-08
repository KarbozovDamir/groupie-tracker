package main

import (
	// "fmt"
	grtrack "grtrack/struct"
	"net/http"
)

//include templates
func main() {
	http.HandleFunc("/main", grtrack.HundlerFirst)
	http.HandleFunc("/bands/", grtrack.HundlerSecond)
	http.ListenAndServe(":8080", nil)
}
