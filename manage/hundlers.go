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

// func MainPageHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		ErrorHandler(w, r, 404)
// 		return
// 	}
// 	if MainErr != nil {
// 		log.Println("Html-file does not exist or it was renamed")
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		return
// 	}
// 	defer r.Body.Close()
// 	if r.Method != http.MethodPost && r.Method != http.MethodGet {
// 		w.Header().Set("Allow", "POST, GET")
// 		ErrorHandler(w, r, 405)
// 		return
// 	}
// 	if r.Method == http.MethodGet {
// 		Main.ExecuteTemplate(w, "mainPage.html", ArtistsVal)
// 		return
// 	}
// }

// //

// func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
// 	// w.WriteHeader(status)
// 	if status == http.StatusBadRequest {
// 		http.Error(w, http.StatusText(status), status)
// 		log.Println("StatusBadRequest")
// 		return
// 	}
// 	if status == http.StatusNotFound {
// 		http.Error(w, http.StatusText(status), status)
// 		log.Println("StatusNotFound")
// 		return
// 	}
// 	if status == http.StatusMethodNotAllowed {
// 		http.Error(w, http.StatusText(status), status)
// 		log.Println("StatusMethodNotAllowed")
// 		return
// 	}
// 	if status == http.StatusInternalServerError {
// 		http.Error(w, http.StatusText(status), status)
// 		log.Println("StatusInternalServerError")
// 		return
// 	}
// }

//

// func ArtistHandler(w http.ResponseWriter, r *http.Request) {
// 	if DifErr != nil {
// 		log.Println("Html-file does not exist or it was renamed")
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		return
// 	}
// 	if r.Method != http.MethodGet {
// 		w.Header().Set("Allow", "GET")
// 		ErrorHandler(w, r, 405)
// 		return
// 	}
// 	defer r.Body.Close()
// 	path := r.URL.Path[8:]
// 	arr, err := strconv.Atoi(path)
// 	if arr <= 0 || arr >= 53 {
// 		log.Println("Index does not exist")
// 		ErrorHandler(w, r, 404)
// 		return
// 	}
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	CreateArtist(arr)
// 	ListGen = append(ListGen, GeneralVal)
// 	Dif.ExecuteTemplate(w, "dif.html", ListGen)
// 	ListGen = []General{}
// }
