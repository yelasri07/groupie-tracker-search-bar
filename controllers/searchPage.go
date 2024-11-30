package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"sync"

	"groupietracker/models"
	"groupietracker/utils"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		renderError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	searchValue := strings.TrimSpace(r.URL.Query().Get("s"))
	if searchValue == "" || len(searchValue) > 100 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var artists models.Data
	err := models.FetchAPI("https://groupietrackers.herokuapp.com/api/artists", &artists.AllArtists)
	if err != nil {
		renderError(w, "Server Error", http.StatusInternalServerError)
		return
	}

	artists.AllArtists[20].Image = "./assets/img/3ib.jpg"

	Search(strings.ToLower(searchValue), &artists)

	artists.Duplicates = utils.RemoveDuplicates(artists.AllArtists)
	artists.HomePage = true

	err = RenderTempalte(w, "./views/index.html", artists, http.StatusOK)
	if err != nil {
		renderError(w, "Server Error", http.StatusInternalServerError)
		return
	}
}

func Search(searchValue string, artists *models.Data) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, artist := range artists.AllArtists {
		wg.Add(1)
		go func(artist models.Artists) {
			defer wg.Done()

			firstSearch := false
			if strings.Contains(strings.ToLower(artist.Name), searchValue) ||
				artist.FirstAlbum == searchValue ||
				strconv.Itoa(artist.CreationDate) == searchValue {
				firstSearch = true
			}

			for _, member := range artist.Members {
				if strings.Contains(strings.ToLower(member), searchValue) {
					firstSearch = true
					break
				}
			}

			if !firstSearch {
				models.FetchAPI(artist.Locations, &artist.Loca)
				for _, localtion := range artist.Loca.Locations {
					if strings.Contains(strings.ToLower(localtion), searchValue) {
						firstSearch = true
						break
					}
				}
			}

			if firstSearch {
				mu.Lock()
				artists.CurrentArtists = append(artists.CurrentArtists, artist)
				mu.Unlock()
			}
		}(artist)
	}

	wg.Wait()
}
