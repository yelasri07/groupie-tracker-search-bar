package models

import (
	"encoding/json"
	"net/http"
)

func FetchAPI(url string, s any) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
		return err
	}

	return nil
}

func GetArtist(url string) (Artists, error) {
	var artist Artists

	err := FetchAPI(url, &artist)
	if err != nil {
		return Artists{}, err
	}

	if artist.ID == 0 {
		return Artists{}, nil
	}

	err = FetchAPI(artist.Locations, &artist.Loca)
	if err != nil {
		return Artists{}, err
	}
	err = FetchAPI(artist.CongertDates, &artist.ConDT)
	if err != nil {
		return Artists{}, err
	}
	err = FetchAPI(artist.Relations, &artist.Rela)
	if err != nil {
		return Artists{}, err
	}

	return artist, nil
}
