package handlers

import (
	"encoding/json"
	"groupie-tracker/models"
	"html/template"
	"net/http"
)

const BaseURL = "https://groupietrackers.herokuapp.com/api"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	resp, err := http.Get(BaseURL + "/artists")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var artists models.ArtistList
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, artists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
