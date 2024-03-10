package types

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Location struct {
	Index []struct {
		Location []string `json:"locations"`
	} `json:"index"`
}

type Relations struct {
	Index []struct {
		Relation
	} `json:"index"`
}

type Artists struct {
	ID           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Locations    []string            `json:"locations2"`
	ConcertDates []string            `json:"-"`
	Relations    map[string][]string `json:"-"`
}
