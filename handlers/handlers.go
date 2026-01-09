package handlers

import (
	"encoding/json"
	"groupie-tracker/models"
	"html/template"
	"net/http"
	"os"
	"strings"
)

const BaseURL = "https://groupietrackers.herokuapp.com/api"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	resp, err := http.Get(BaseURL + "/artists")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer resp.Body.Close()

	var artists models.ArtistList
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	for i := range artists {
		artists[i].FirstAlbum = strings.ReplaceAll(artists[i].FirstAlbum, "-", "/")
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = tmpl.Execute(w, artists)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.NotFound(w, r)
		return
	}

	// Lecture du fichier de config secret
	configFile, err := os.Open("config.json")
	if err != nil {
		http.Error(w, "Erreur: Impossible de lire config.json", 500)
		return
	}
	defer configFile.Close()

	var config models.Config
	if err := json.NewDecoder(configFile).Decode(&config); err != nil {
		http.Error(w, "Erreur JSON Config", 500)
		return
	}

	respArtist, err := http.Get(BaseURL + "/artists/" + id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer respArtist.Body.Close()

	var artist models.Artist
	if err := json.NewDecoder(respArtist.Body).Decode(&artist); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	artist.FirstAlbum = strings.ReplaceAll(artist.FirstAlbum, "-", "/")

	respRel, err := http.Get(BaseURL + "/relation/" + id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer respRel.Body.Close()

	var relation models.Relation
	if err := json.NewDecoder(respRel.Body).Decode(&relation); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := models.PageData{
		Artist:   artist,
		Relation: relation,
		ApiKey:   config.YoutubeKey,
	}

	tmpl, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tmpl.Execute(w, data)
}
