package grtrack

import (
	"io/ioutil"
	"log"
	"net/http"
)

//using golang’s http module to get the response:
func getApi() {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists") //download data from data API

	if err != nil {
		log.Fatal(err)
	}

	// reading the http body into a byte array for parsing/processing (using golangs ioutil library)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
}

//API artists
//https:groupietrackers.herokuapp.com/api/artists

//API relation
//https:groupietrackers.herokuapp.com/api/relation
