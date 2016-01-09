package aura

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ArtistsInterface interface {
	GetAllArtists(limit int, cont string) ([]Artist, error)
	GetArtistByID(artistID string) (Artist, error)
}

type ArtistsFeature struct {
	ArtistsInterface
	Feature
}

func (f *ArtistsFeature) Routes() Routes {
	return Routes{
		Route{
			"AllArtists",
			"GET", "/artists", f.AllArtists,
		},
		Route{
			"Artist",
			"GET", "/artists/{artistID}", f.Artist,
		},
		Route{
			"ArtistImages",
			"GET", "/artists/{artistID}/images", f.ArtistImages,
		},
		Route{
			"ArtistImage",
			"GET", "/artists/{artistID}/images/{imageID}", f.ArtistImage,
		},
	}
}

func (f *ArtistsFeature) AllArtists(w http.ResponseWriter, r *http.Request) {
	artists, err := f.GetAllArtists(-1, "")
	if err != nil {
		log.Fatal(err)
	}

	f.sendJSON(w, artists)
}

func (f *ArtistsFeature) Artist(w http.ResponseWriter, r *http.Request) {
	artistID := mux.Vars(r)["artistID"]

	artist, err := f.GetArtistByID(artistID)
	if err != nil {
		log.Fatal(err)
	}

	f.sendJSON(w, artist)
}

func (f *ArtistsFeature) ArtistImages(w http.ResponseWriter, r *http.Request) {
}

func (f *ArtistsFeature) ArtistImage(w http.ResponseWriter, r *http.Request) {
}
