package functions

import (
	"strconv"
	"strings"

	"groupietracker/database"
)

func Search(searchValue string) (database.Data, error) {
	var ArtistsData database.Data

	cachedArtistsData, ok := Cache.Load("Artists")
	if ok {
		ArtistsData.AllArtists = cachedArtistsData.([]database.Artists)
	} else {
		err := StoreDataCache(&ArtistsData.AllArtists)
		if err != nil {
			return database.Data{}, err
		}
		Cache.Store("Artists", ArtistsData.AllArtists)
	}

	var firstSearch bool
	for _, artist := range ArtistsData.AllArtists {
		firstSearch = false
		if strings.Contains(strings.ToLower(artist.Name), searchValue) ||
			artist.FirstAlbum == searchValue ||
			strconv.Itoa(artist.CreationDate) == searchValue {
			firstSearch = true
		}

		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), searchValue) {
				firstSearch = true
				continue
			}
		}

		for _, localtion := range artist.Loca.Locations {
			if strings.Contains(strings.ToLower(localtion), searchValue) {
				firstSearch = true
				continue
			}
		}

		if firstSearch {
			ArtistsData.CurrentArtists = append(ArtistsData.CurrentArtists, artist)
		}

	}

	return ArtistsData, nil
}
