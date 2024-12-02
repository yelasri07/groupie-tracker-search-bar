package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"sync"

	"groupietracker/models"
	"groupietracker/utils"
)

// search page to view artists after search
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

	Search(strings.ToLower(searchValue), &artists)

	artists.Duplicates = utils.RemoveDuplicates(artists.AllArtists)
	artists.HomePage = true

	err := RenderTempalte(w, "./views/index.html", artists, http.StatusOK)
	if err != nil {
		renderError(w, "Server Error", http.StatusInternalServerError)
		return
	}
}

func Search(searchValue string, artists *models.Data) {
	var locations models.IndexLocations
	models.FetchAPI("https://groupietrackers.herokuapp.com/api/artists", &artists.AllArtists)
	models.FetchAPI("https://groupietrackers.herokuapp.com/api/locations", &locations)

	artists.AllArtists[20].Image = "./assets/img/3ib.jpg"
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
				idLocation := strings.Split(artist.Locations, "/")[5]
				for _, location := range locations.Index {
					if idLocation == strconv.Itoa(location.ID) {
						for _, loca := range location.Locations {
							if strings.Contains(strings.ToLower(loca), searchValue) {
								firstSearch = true
								break
							}
						}
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
