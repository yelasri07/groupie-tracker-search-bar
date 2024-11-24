package controllers

import (
	"net/http"
	"strings"

	"groupietracker/functions"
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

	ArtistsData, err := functions.Search(strings.ToLower(searchValue))
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
