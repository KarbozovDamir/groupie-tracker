package grtrack

import (
	"fmt"
	"net/http"
)

func HundlerFirst(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "fun")
}

func HundlerSecond(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "fun2")
}
func ErrorHandler(w http.ResponseWriter, status int) {
	switch status {
	case 400:
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	case 404:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	case 405:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	default:
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
