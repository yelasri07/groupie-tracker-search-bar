package controllers

import (
	"html/template"
	"net/http"

	"groupietracker/database"
)

func RenderTempalte(w http.ResponseWriter, url string, data any, status int) error {
	tmpl, err := template.ParseFiles(url)
	if err != nil {
		return err
	}

	w.WriteHeader(status)
	err = tmpl.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}

func renderError(w http.ResponseWriter, typeError string, status int) {
	e := database.ErrorPage{Status: status, Type: typeError}
	RenderTempalte(w, "templates/error.html", e, status)
}
