package aura

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type TracksInterface interface {
	GetAllTracks(limit int, cont string) (TrackList, error)
	GetTrackByID(trackID string) (Track, error)
}

type TracksFeature struct {
	Feature
	TracksInterface
}

func (f *TracksFeature) Routes() Routes {
	return Routes{
		Route{
			"AllTracks",
			"GET", "/tracks", f.AllTracks,
		},
		Route{
			"Track",
			"GET", "/tracks/{trackID}", f.Track,
		},
		Route{
			"TrackAudio",
			"GET", "/tracks/{trackID}/audio", f.TrackAudio,
		},
		Route{
			"TrackImages",
			"GET", "/tracks/{trackID}/images", f.TrackImages,
		},
		Route{
			"TrackImage",
			"GET", "/tracks/{trackID}/images/{imageID}", f.TrackImage,
		},
	}
}

type AllTracksResponse struct {
	Tracks TrackList `json:"tracks"`
}

func (f *TracksFeature) AllTracks(w http.ResponseWriter, r *http.Request) {
	tracks, err := f.GetAllTracks(-1, "")
	if err != nil {
		log.Fatal(err)
	}

	f.sendJSON(w, AllTracksResponse{
		Tracks: tracks,
	})
}

func (f *TracksFeature) Track(w http.ResponseWriter, r *http.Request) {
	trackID := mux.Vars(r)["trackID"]

	track, err := f.GetTrackByID(trackID)
	if err != nil {
		log.Fatal(err)
	}

	f.sendJSON(w, track)
}

func (f *TracksFeature) TrackAudio(w http.ResponseWriter, r *http.Request) {
}

func (f *TracksFeature) TrackImages(w http.ResponseWriter, r *http.Request) {
}

func (f *TracksFeature) TrackImage(w http.ResponseWriter, r *http.Request) {
}
