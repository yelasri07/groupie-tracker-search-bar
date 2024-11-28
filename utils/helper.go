package utils

import (
	"strconv"

	"groupietracker/models"
)

func RemoveDuplicates(artists []models.Artists, RmDup map[string]string) map[string]string {
	for _, artist := range artists {
		RmDup[artist.Name] = "artist/band"
		RmDup[artist.FirstAlbum] = "First Album"
		RmDup[strconv.Itoa(artist.CreationDate)] = "Creation Date"
		for _, member := range artist.Members {
			RmDup[member] = "member"
		}

		for _, location := range artist.Loca.Locations {
			RmDup[location] = "location"
		}
	}

	return RmDup
}
