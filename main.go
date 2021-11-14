package main

import (
	"fmt"
	grtrack "grtrack/struct"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/main", grtrack.HundlerFirst)
	http.HandleFunc("/bands/", grtrack.HundlerSecond)
	fmt.Printf("The server is running on this address: http://localhost:8081\n")
	http.ListenAndServe(":8080", nil)
	// log.Fatal(ListenAndServe(":8080", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "webApiBd!")
}
