package controllers

import (
	"html/template"
	"net/http"

	"groupietracker/models"
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
	e := models.ErrorPage{Status: status, Type: typeError}
	RenderTempalte(w, "views/error.html", e, status)
}
