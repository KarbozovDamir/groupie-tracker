package grtrack

import (
	"encoding/json"
	"fmt"
)

// func HundlerFirst(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println(w, "fun")
// }

// func HundlerSecond(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println(w, "fun2")
// }
// func ErrorHandler(w http.ResponseWriter, status int) {
// 	switch status {
// 	case 400:
// 		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
// 	case 404:
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 	case 405:
// 		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
// 	default:
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 	}
// }

func getBody(body []byte) (*body, error) {
	var s = new(artistData)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}

// }
// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World!")
// }

// //jsonHandler returns http respone in JSON format.
// func jsonHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	user := User{Id: 1,
// 		Name:  "John Doe",
// 		Email: "johndoe@gmail.com",
// 		Phone: "000099999"}
// 	json.NewEncoder(w).Encode(user)
// }

// //templateHandler renders a template and returns as http response.
// func templateHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	t, err := template.ParseFiles("template.html")
// 	if err != nil {
// 		fmt.Fprintf(w, "Unable to load template")
// 	}

// 	user := User{Id: 1,
// 		Name:  "John Doe",
// 		Email: "johndoe@gmail.com",
// 		Phone: "000099999"}
// 	t.Execute(w, user)
// }
