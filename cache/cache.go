package cache

import (
	"sync"

	"groupietracker/models"
)

var Cache sync.Map

func SaveToCache(artists *[]models.Artists, key string) error {
	err := models.FetchAPI("https://groupietrackers.herokuapp.com/api/artists", &artists)
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
			errChann <- models.GetForeignData(&(*artists)[index])
		}(i)
	}

	wg.Wait()
	close(errChann)

	for err := range errChann {
		if err != nil {
			return err
		}
	}

	Cache.Store(key, *artists)

	return nil
}

func GetFromCache(key string) (any, bool) {
	return Cache.Load(key)
}

func RemoveFromCache(key string) {
	Cache.Delete(key)
}
