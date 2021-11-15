package main

import (
	"fmt"
	"net/http"
)

// func main() {
// 	// http.HandleFunc("/", handler)
// 	// http.HandleFunc("/main", grtrack.HundlerFirst)
// 	// http.HandleFunc("/bands/", grtrack.HundlerSecond)
// 	fmt.Printf("The server is running on this address: http://localhost:8080")
// 	http.ListenAndServe(":8080", nil)
// 	// log.Fatal(ListenAndServe(":8080", nil))

// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "webApiBd!")
// }

// func main() {
// 	http.HandleFunc("/", handler)
// 	http.HandleFunc("/json", jsonHandler)
// 	http.HandleFunc("/template", templateHandler)
// 	http.ListenAndServe(":8081", nil)
// }

//using golang’s http module to get the response:
func main() {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api") //download data from data API

	if err != nil {
		log.Fatal(err)
	}

	// reading the http body into a byte array for parsing/processing (using golangs ioutil library)
	body, err := ioutil.ReadAll(res.Body)
if err != nil {
	panic(err.Error())
	
    s, err := http.Get([]byte(body))
}