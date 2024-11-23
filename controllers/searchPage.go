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

	searchValue := r.URL.Query().Get("s")
	if searchValue == "" {
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

func Search(searchValue string) (data, error) {
	var ArtistsData data

	cachedArtistsData, ok := cache.Load("Artists")
	if ok {
		ArtistsData.AllArtists = cachedArtistsData.([]database.Artists)
	} else {
		err := storeDataCache(&ArtistsData.AllArtists)
		if err != nil {
			return data{}, err
		}
		cache.Store("Artists", ArtistsData.AllArtists)
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
