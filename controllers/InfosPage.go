package controllers

import (
	"net/http"
	"strconv"

	"groupietracker/database"
)

func InfosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.Query().Get("id")

		if id != "" {

			ID, err := strconv.Atoi(id)
			if err != nil {
				e := database.ErrorPage{Status: 404, Type: "User not found"}
				RenderTempalte(w, "templates/error.html", e, http.StatusNotFound)
				return
			}

			id = strconv.Itoa(ID)

			var artist database.Artists

			err = database.FetchAPI("https://groupietrackers.herokuapp.com/api/artists/"+id, &artist)
			if err != nil {
				e := database.ErrorPage{Status: 500, Type: "Server Error"}
				RenderTempalte(w, "templates/error.html", e, http.StatusInternalServerError)
				return
			}

			if artist.ID == 0 {
				e := database.ErrorPage{Status: 404, Type: "User not found"}
				RenderTempalte(w, "templates/error.html", e, http.StatusNotFound)
				return
			}

			err = database.GetForeignData(&artist)
			if err != nil {
				e := database.ErrorPage{Status: 500, Type: "Server Error"}
				RenderTempalte(w, "templates/error.html", e, http.StatusInternalServerError)
				return
			}

			if artist.ID == 21 {
				artist.Image = "./assets/img/3ib.jpg"
			}

			err = RenderTempalte(w, "./templates/infos.html", artist, http.StatusOK)
			if err != nil {
				e := database.ErrorPage{Status: 500, Type: "Server Error"}
				RenderTempalte(w, "templates/error.html", e, http.StatusInternalServerError)
				return
			}
		} else {
			e := database.ErrorPage{Status: 400, Type: "Bad Request"}
			RenderTempalte(w, "templates/error.html", e, http.StatusBadRequest)
			return
		}
	} else {
		e := database.ErrorPage{Status: 405, Type: "Method Not Allowed"}
		RenderTempalte(w, "templates/error.html", e, http.StatusMethodNotAllowed)
		return
	}
}
