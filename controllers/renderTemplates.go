package controllers

import (
	"html/template"
	"net/http"
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
