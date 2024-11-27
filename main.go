package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"groupietracker/routers"
)

const port string = ":8082"

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Please enter only the program name.")
		return
	}

	routers.Routers()
	fmt.Println("http://localhost" + port + "/")
	log.Fatal(http.ListenAndServe(port, nil))
}
