package grtrack



//using golang’s http module to get the response:
func getApi() {
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