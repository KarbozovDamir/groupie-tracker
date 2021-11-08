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
