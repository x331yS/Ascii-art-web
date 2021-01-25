package pkg

import (
	"log"
	"net/http"
	"html/template"
)

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	t, _ := template.ParseFiles("templates/500.html")
	err := t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func BadRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	t, err := template.ParseFiles("templates/400.html")
	if err != nil {
		InternalServerError(w, r)
	}
	t.Execute(w, nil)
}

func StatusNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	t, err := template.ParseFiles("templates/404.html")
	if err != nil {
		InternalServerError(w, r)
		return
	}
	t.Execute(w, nil)
}