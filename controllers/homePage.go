package controllers

import (
	"net/http"
	"sync"

	"groupietracker/database"
)

var cache sync.Map

type data struct{
	AllArtists []database.Artists
	CurrentArtists []database.Artists
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		renderError(w, "Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		renderError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var ArtistsData data

	if cachedData, ok := cache.Load("Artists"); ok {
		ArtistsData.AllArtists = cachedData.([]database.Artists)
	} else {
		err := storeDataCache(&ArtistsData.AllArtists)
		if err != nil {
			renderError(w, "Server Error", http.StatusInternalServerError)
			return
		}
		cache.Store("Artists", ArtistsData.AllArtists)
	}

	ArtistsData.CurrentArtists = ArtistsData.AllArtists

	err := RenderTempalte(w, "./templates/index.html", ArtistsData, http.StatusOK)
	if err != nil {
		renderError(w, "Server Error", http.StatusInternalServerError)
		return
	}
}

func storeDataCache(artists *[]database.Artists) error {
	err := database.FetchAPI("https://groupietrackers.herokuapp.com/api/artists", &artists)
	if err != nil {
		return err
	}

	(*artists)[20].Image = "assets/img/3ib.jpg"

	var wg sync.WaitGroup
	errChann := make(chan error, len(*artists))

	for i := 0; i < len(*artists); i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			errChann <- database.GetForeignData(&(*artists)[index])
		}(i)
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
