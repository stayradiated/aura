package aura

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ArtistsInterface interface {
	FilterArtists(params map[string]string, limit int, cont string) (ArtistList, error)
	ArtistWithID(artistID string) (Artist, error)
}

type ArtistsFeature struct {
	ArtistsInterface
	Feature
}

func (f *ArtistsFeature) Routes() Routes {
	return Routes{
		Route{
			"AllArtists",
			"GET", "/artists", f.getArtists,
		},
		Route{
			"Artist",
			"GET", "/artists/{artistID}", f.getArtistWithID,
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

type ArtistsResponse struct {
	Artists []Artist `json:"artists"`
}

func (f *ArtistsFeature) getArtists(w http.ResponseWriter, r *http.Request) {
	params := f.getQueryParams(r)

	artists, err := f.FilterArtists(params, -1, "")
	if err != nil {
		log.Fatal(err)
	}

	f.sendJSON(w, ArtistsResponse{
		Artists: artists,
	})
}

func (f *ArtistsFeature) getArtistWithID(w http.ResponseWriter, r *http.Request) {
	artistID := mux.Vars(r)["artistID"]

	artist, err := f.ArtistWithID(artistID)
	if err != nil {
		log.Fatal(err)
	}

	f.sendJSON(w, ArtistsResponse{
		Artists: []Artist{artist},
	})
}

func (f *ArtistsFeature) ArtistImages(w http.ResponseWriter, r *http.Request) {
}

func (f *ArtistsFeature) ArtistImage(w http.ResponseWriter, r *http.Request) {
}
