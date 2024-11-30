package controllers

import (
	"net/http"

	"groupietracker/models"
	"groupietracker/utils"
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

	var artists models.Data

	err := models.FetchAPI("https://groupietrackers.herokuapp.com/api/artists", &artists.AllArtists)
	if err != nil {
		renderError(w, "Server Error", http.StatusInternalServerError)
	}

	artists.AllArtists[20].Image = "./assets/img/3ib.jpg"

	artists.CurrentArtists = artists.AllArtists
	artists.RmDup = utils.RemoveDuplicates(artists.AllArtists)

	err = RenderTempalte(w, "./views/index.html", artists, http.StatusOK)
	if err != nil {
		renderError(w, "Server Error", http.StatusInternalServerError)
		return
	}
}
