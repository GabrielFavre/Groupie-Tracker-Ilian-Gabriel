package handlers

import (
	"groupie-tracker/api"
	"groupie-tracker/models"
	"html/template"
	"net/http"
	"strconv"
)

const BaseURL = "https://groupietrackers.herokuapp.com/api"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	var artists models.ArtistList
	err := api.FetchData(BaseURL+"/artists", &artists)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, artists)
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	var artist models.Artist
	err = api.FetchData(BaseURL+"/artists/"+idStr, &artist)
	if err != nil {
		http.Error(w, "Error fetching artist", http.StatusInternalServerError)
		return
	}

	var relation models.Relation
	err = api.FetchData(BaseURL+"/relation/"+idStr, &relation)
	if err != nil {
		http.Error(w, "Error fetching relation", http.StatusInternalServerError)
		return
	}

	data := struct {
		Artist   models.Artist
		Relation models.Relation
	}{
		Artist:   artist,
		Relation: relation,
	}

	tmpl, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}
