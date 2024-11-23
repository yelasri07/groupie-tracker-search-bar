package controllers

import (
	"net/http"
	"os"
)

func CssHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		renderError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	file, err := os.Stat(r.URL.Path[1:])
	if err != nil || file.IsDir() {
		renderError(w, "Page Not Found", http.StatusNotFound)
		return
	}

	fs := http.Dir("assets/css/")
	http.StripPrefix("/assets/css/", http.FileServer(fs)).ServeHTTP(w, r)
}

func ImagesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		renderError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	file, err := os.Stat(r.URL.Path[1:])
	if err != nil || file.IsDir() {
		renderError(w, "Page Not Found", http.StatusNotFound)
		return
	}

	fs := http.Dir("assets/img/")
	http.StripPrefix("/assets/img/", http.FileServer(fs)).ServeHTTP(w, r)
}
