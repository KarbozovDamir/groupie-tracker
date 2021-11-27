package main

import (
	"fmt"
	"log"
	"net/http"
)

type b int

func main() {
	mux := http.NewServeMux()

	// b := beka
	mux.HandleFunc("/", handler) //home page

	// mux.HandleFunc("/band/", grtrack.Band) // one group or one artist
	fmt.Printf("The server is running on this address: http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", mux))

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "webApiBd!")

}

// func main() {
// 	http.HandleFunc("/", handler)
// 	http.HandleFunc("/json", jsonHandler)
// 	http.HandleFunc("/template", templateHandler)
// 	http.ListenAndServe(":8081", nil)
// }
