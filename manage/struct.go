package grtrack

type artistData struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

// dates of location
type relation struct {
	ID       int                 `json:"id"`
	Concerts map[string][]string `json:"datesLocations"`
}

var allData []artistData
