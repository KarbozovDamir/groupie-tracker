package grtrack

import (
	"fmt"
	"net/http"
)

func Hundler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "fun")
}

func HundlerBand(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "fun2")
}

func ErrorHandler(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
	return
}

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
