package controllers

import (
	"net/http"
	"sync"

	"groupietracker/database"
)

var cache sync.Map

func storeDataCache(artists *[]database.Artists) error {
	err := database.FetchAPI("https://groupietrackers.herokuapp.com/api/artists", &artists)
	if err != nil {
		return err
	}

	(*artists)[20].Image = "assets/img/3ib.jpg"

	var wg sync.WaitGroup
	errChann := make(chan error, len(*artists))

	for i := 0; i < len(*artists); i++ {
		artist := (*artists)[i]
		wg.Add(1)
		go func(artist database.Artists) {
			defer wg.Done()
			errChann <- database.GetForeignData(&artist)
		}(artist)
	}

	wg.Wait()
	close(errChann)

	for err := range errChann {
		if err != nil {
			return err
		}
	}

	return nil
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		if r.Method == http.MethodGet {
			var artists []database.Artists
			if cachedData, ok := cache.Load("Artists"); ok {
				artists = cachedData.([]database.Artists)
			} else {
				err := storeDataCache(&artists)
				if err != nil {
					e := database.ErrorPage{Status: 500, Type: "Server Error"}
					RenderTempalte(w, "templates/error.html", e, http.StatusInternalServerError)
					return
				}
				cache.Store("Artists", artists)
			}

			err := RenderTempalte(w, "./templates/index.html", artists, http.StatusOK)
			if err != nil {
				e := database.ErrorPage{Status: 500, Type: "Server Error"}
				RenderTempalte(w, "templates/error.html", e, http.StatusInternalServerError)
				return
			}

		} else {
			e := database.ErrorPage{Status: 405, Type: "Method Not Allowed"}
			RenderTempalte(w, "templates/error.html", e, http.StatusMethodNotAllowed)
			return
		}
	default:
		e := database.ErrorPage{Status: 404, Type: "Page Not Found"}
		RenderTempalte(w, "templates/error.html", e, http.StatusNotFound)
		return
	}
}
