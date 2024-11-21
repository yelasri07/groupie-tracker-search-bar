package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"groupietracker/controllers"
)

const port string = ":8082"

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Please enter only the program name.")
		return
	}

	http.HandleFunc("/assets/img/", controllers.ImagesHandler)
	http.HandleFunc("/assets/css/", controllers.CssHandler)
	http.HandleFunc("/", controllers.IndexHandler)
	http.HandleFunc("/infos", controllers.InfosHandler)
	http.HandleFunc("/sch", controllers.SearchHandler)
	fmt.Println("http://localhost" + port + "/")
	log.Fatal(http.ListenAndServe(port, nil))
}
