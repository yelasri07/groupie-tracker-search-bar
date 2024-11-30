package models

type Data struct {
	AllArtists     []Artists
	CurrentArtists []Artists
	Duplicates map[string]string
	HomePage bool
}

type Artists struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	Loca         Locations
	CongertDates string `json:"concertDates"`
	ConDT        Dates
	Relations    string `json:"relations"`
	Rela         Relations
}

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type ErrorPage struct {
	Status int
	Type   string
}
