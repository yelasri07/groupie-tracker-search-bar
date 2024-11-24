package functions

import (
	"sync"

	"groupietracker/database"
)

var Cache sync.Map

func StoreDataCache(artists *[]database.Artists) error {
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
