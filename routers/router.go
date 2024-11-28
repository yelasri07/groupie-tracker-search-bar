package routers

import (
	"net/http"

	"groupietracker/controllers"
)

func Routers() {
	http.HandleFunc("/assets/", controllers.AssetsHandler)
	http.HandleFunc("/", controllers.IndexHandler)
	http.HandleFunc("/infos", controllers.InfosHandler)
	http.HandleFunc("/sch",controllers.SearchHandler)
}
