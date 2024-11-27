package routers

import (
	"net/http"

	"groupietracker/controllers"
)

func Routers() {
	http.HandleFunc("/assets/img/", controllers.ImagesHandler)
	http.HandleFunc("/assets/css/", controllers.CssHandler)
	http.HandleFunc("/", controllers.IndexHandler)
	http.HandleFunc("/infos", controllers.InfosHandler)
}
