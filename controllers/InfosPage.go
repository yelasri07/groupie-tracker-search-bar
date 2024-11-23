package controllers

import (
	"net/http"
	"strconv"

	"groupietracker/database"
)

func InfosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		renderError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		renderError(w, "Bad Request", http.StatusBadRequest)
		return
	}

	ID, err := strconv.Atoi(id)
	if err != nil {
		renderError(w, "Artist Not Found", http.StatusNotFound)
		return
	}

	id = strconv.Itoa(ID)

	artist, err := database.GetArtist("https://groupietrackers.herokuapp.com/api/artists/" + id)
	if err != nil {
		renderError(w, "Server Error", http.StatusInternalServerError)
		return
	}

	if artist.ID == 0 {
		renderError(w, "Artist Not Found", http.StatusNotFound)
		return
	}

	if artist.ID == 21 {
		artist.Image = "./assets/img/3ib.jpg"
	}

	err = RenderTempalte(w, "./templates/infos.html", artist, http.StatusOK)
	if err != nil {
		renderError(w, "Server Error", http.StatusInternalServerError)
		return
	}
}
