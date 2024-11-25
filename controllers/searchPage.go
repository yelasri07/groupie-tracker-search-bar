package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"groupietracker/database"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		renderError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	searchValue := strings.TrimSpace(r.URL.Query().Get("s"))
	if searchValue == "" || len(searchValue) > 50 {
		renderError(w, "Bad Request", http.StatusBadRequest)
		return
	}

	ArtistsData, err := Search(strings.ToLower(searchValue))
	if err != nil {
		renderError(w, "Server Error", http.StatusInternalServerError)
		return
	}

	err = RenderTempalte(w, "./templates/index.html", ArtistsData, http.StatusOK)
	if err != nil {
		renderError(w, "Server Error", http.StatusInternalServerError)
		return
	}
}

func Search(searchValue string) (database.Data, error) {
	var ArtistsData database.Data

	cachedArtistsData, ok := database.Cache.Load("Artists")
	if ok {
		ArtistsData.AllArtists = cachedArtistsData.([]database.Artists)
	} else {
		err := database.StoreDataCache(&ArtistsData.AllArtists)
		if err != nil {
			return database.Data{}, err
		}
		database.Cache.Store("Artists", ArtistsData.AllArtists)
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
				break
			}
		}

		for _, localtion := range artist.Loca.Locations {
			if strings.Contains(strings.ToLower(localtion), searchValue) {
				firstSearch = true
				break
			}
		}

		if firstSearch {
			ArtistsData.CurrentArtists = append(ArtistsData.CurrentArtists, artist)
		}

	}

	return ArtistsData, nil
}
