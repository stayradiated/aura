package aura

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ArtistsInterface interface {
	FilterArtists(params map[string]string, include Include, limit int, cont string) (Entities, error)
	ArtistWithID(artistID string, include Include) (Entities, error)
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
			"GET", "/artists/{artistID}/images", f.getArtistImages,
		},
		Route{
			"ArtistImage",
			"GET", "/artists/{artistID}/images/{imageID}", f.getArtistImage,
		},
	}
}

type ArtistsResponse struct {
	Artists []Artist `json:"artists"`
}

func (f *ArtistsFeature) getArtists(w http.ResponseWriter, r *http.Request) {
	params := f.getQueryParams(r)
	include := f.getInclude(&params)
	limit := f.getLimit(&params)

	entities, err := f.FilterArtists(params, include, limit, "")
	if err != nil {
		log.Fatal(err)
	}

	f.sendJSON(w, Response{
		Result:   entities.Artists.IDs(),
		Entities: entities.asMap(),
	})
}

func (f *ArtistsFeature) getArtistWithID(w http.ResponseWriter, r *http.Request) {
	artistID := mux.Vars(r)["artistID"]
	params := f.getQueryParams(r)
	include := f.getInclude(&params)

	entities, err := f.ArtistWithID(artistID, include)
	if err != nil {
		log.Fatal(err)
	}

	f.sendJSON(w, Response{
		Result:   entities.Artists.IDs(),
		Entities: entities.asMap(),
	})
}

func (f *ArtistsFeature) getArtistImages(w http.ResponseWriter, r *http.Request) {
}

func (f *ArtistsFeature) getArtistImage(w http.ResponseWriter, r *http.Request) {
}
