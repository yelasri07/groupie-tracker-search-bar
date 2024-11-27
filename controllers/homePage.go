package controllers

import (
	"net/http"

	"groupietracker/cache"
	"groupietracker/database"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		renderError(w, "Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		renderError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var ArtistsData database.Data

	if cachedData, ok := cache.GetFromCache("Artists"); ok {
		ArtistsData.AllArtists = cachedData.([]database.Artists)
	} else {
		err := cache.SaveToCache("Artists")
		if err != nil {
			renderError(w, "Server Error", http.StatusInternalServerError)
			return
		}
		database.Cache.Store("Artists", ArtistsData.AllArtists)
	}

	ArtistsData.CurrentArtists = ArtistsData.AllArtists

	err := RenderTempalte(w, "./templates/index.html", ArtistsData, http.StatusOK)
	if err != nil {
		renderError(w, "Server Error", http.StatusInternalServerError)
		return
	}
}
