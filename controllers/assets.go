package controllers

import (
	"net/http"
	"os"
)

// Hander to serve css and images
func AssetsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		renderError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	file, err := os.Stat(r.URL.Path[1:])
	if err != nil || file.IsDir() {
		renderError(w, "Page Not Found", http.StatusNotFound)
		return
	}

	fs := http.Dir("assets/")
	http.StripPrefix("/assets/", http.FileServer(fs)).ServeHTTP(w, r)
}
