package controllers

import (
	"net/http"
	"strings"

	"groupietracker/database"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		searchValue := r.URL.Query().Get("s")

		if searchValue != "" {
			var data []database.Artists

			err := Search(strings.ToLower(searchValue), &data)
			if err != nil {
				e := database.ErrorPage{Status: 500, Type: "Server Error"}
				RenderTempalte(w, "templates/error.html", e, http.StatusInternalServerError)
				return
			}

			err = RenderTempalte(w, "./templates/index.html", data, http.StatusOK)
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

func Search(searchValue string, data *[]database.Artists) error {
	var artists []database.Artists
	err := database.FetchAPI("https://groupietrackers.herokuapp.com/api/artists", &artists)
	if err != nil {
		return err
	}

	errChann := make(chan error, len(artists))
	doneGoroutines := make(chan bool)

	for i := 0; i < len(artists); i++ {
		artist := artists[i]
		go func(artist database.Artists) {
			select {
			case errChann <- database.GetForeignData(&artist):
			case <-doneGoroutines:
				return
			}
		}(artist)
	}

	for i := 0; i < len(artists); i++ {
		if err := <-errChann; err != nil {
			close(doneGoroutines)
			return err
		}
	}

	for _, artist := range artists {
		if strings.HasPrefix(strings.ToLower(artist.Name), searchValue) {
			*data = append(*data, artist)
		}

		for _, member := range artist.Members {
			if strings.HasPrefix(strings.ToLower(member), searchValue) {
				*data = append(*data, artist)
			}
		}

	}

	return nil
}
