package controllers

import (
	"net/http"
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

	data, err := Search(strings.ToLower(searchValue))
	if err != nil {
		renderError(w, "Server Error", http.StatusInternalServerError)
		return
	}

	err = RenderTempalte(w, "./templates/index.html", data, http.StatusOK)
	if err != nil {
		renderError(w, "Server Error", http.StatusInternalServerError)
		return
	}
}

func Search(searchValue string) ([]database.Artists, error) {
	var artists []database.Artists
	
	cachedData, ok := cache.Load("Artists")
	if ok {
		artists = cachedData.([]database.Artists)
	} else {
		err := storeDataCache(&artists)
		if err != nil {
			return nil, err
		}
		cache.Store("Artists", artists)
	}

	var data []database.Artists

	for _, artist := range artists {
		if strings.HasPrefix(strings.ToLower(artist.Name), searchValue) {
			data = append(data, artist)
		}

		for _, member := range artist.Members {
			if strings.HasPrefix(strings.ToLower(member), searchValue) {
				data = append(data, artist)
			}
		}

	}

	return data, nil
}
