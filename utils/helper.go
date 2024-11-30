package utils

import (
	"strconv"

	"groupietracker/models"
)

func RemoveDuplicates(artists []models.Artists) map[string]string {
	var locations []models.Locations
	models.FetchAPI("https://groupietrackers.herokuapp.com/api/locations",&locations)

	duplicates := make(map[string]string)
	for _, artist := range artists {
		duplicates[artist.Name] = "artist/band"
		duplicates[artist.FirstAlbum] = "first album"
		duplicates[strconv.Itoa(artist.CreationDate)] = "creation date"
		for _, member := range artist.Members {
			duplicates[member] = "member"
		}
	}

	for _, location := range locations{
		for _, loca := range location.Locations {
			duplicates[loca] = "location"
		}
	}

	return duplicates
}
