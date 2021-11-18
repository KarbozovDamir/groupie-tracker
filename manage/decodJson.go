package grtrack

import (
	"encoding/json"
	"fmt"
)

//get json and create "artists"

func getBody(body []byte) (*body, error) {
	var s = new(artistData)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}
