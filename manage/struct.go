package grtrack

// type artistData struct {
// 	ID           int      `json:"id"`
// 	Image        string   `json:"image"`
// 	Name         string   `json:"name"`
// 	Members      []string `json:"members"`
// 	CreationDate int      `json:"creationDate"`
// 	FirstAlbum   string   `json:"firstAlbum"`
// }

// // dates of location
// type relation struct {
// 	ID       int                 `json:"id"`
// 	Concerts map[string][]string `json:"datesLocations"`
// }

// var allData []artistData

// type General struct {
// 	Artists   *Artists
// 	Locations struct {
// 		ID        int      "json:\"id\""
// 		Locations []string "json:\"locations\""
// 		Dates     string   "json:\"dates\""
// 	}
// 	ConcertDates struct {
// 		ID    int      `json:"id"`
// 		Dates []string `json:"dates"`
// 	}
// 	Relations struct {
// 		ID             int                 `json:"id"`
// 		DatesLocations map[string][]string `json:"datesLocations"`
// 	}
// }

type Artists []struct {
	ID           int      `json:"id`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	// Location       string   `json:"locations"`
	// ConcertDates   string   `json:"concertDates"`
	DatesLocations map[string][]string
}

type Locations struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

type ConcertDates struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

type Relations struct {
	Index []struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}
