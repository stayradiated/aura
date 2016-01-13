package aura

import (
	"log"
	"net/http"
)

type TracksInterface interface {
	FilterTracks(params map[string]string, include map[string]bool, limit int, cont string) (TrackList, *Included, error)
	TrackWithID(trackID string, include map[string]bool) (Track, *Included, error)
}

type TracksFeature struct {
	Feature
	TracksInterface
}

func (f *TracksFeature) Routes() Routes {
	return Routes{
		Route{
			"GetTracks",
			"GET", "/tracks", f.getTracks,
		},
		Route{
			"GetTrackWithID",
			"GET", "/tracks/{trackID}", f.getTrackWithID,
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

type TracksResponse struct {
	Tracks TrackList `json:"tracks"`
	Links  *Included `json:"links,omitempty"`
}

func (f *TracksFeature) getTracks(w http.ResponseWriter, r *http.Request) {
	params := f.getQueryParams(r)
	include := f.getInclude(&params)
	limit := f.getLimit(&params)

	tracks, included, err := f.FilterTracks(params, include, limit, "")
	if err != nil {
		log.Fatal(err)
	}

	f.sendJSON(w, TracksResponse{
		Tracks: tracks,
		Links:  included,
	})
}

func (f *TracksFeature) getTrackWithID(w http.ResponseWriter, r *http.Request) {
	trackID := f.getVar(r, "trackID")
	params := f.getQueryParams(r)
	include := f.getInclude(&params)

	track, included, err := f.TrackWithID(trackID, include)
	if err != nil {
		log.Fatal(err)
	}

	f.sendJSON(w, TracksResponse{
		Tracks: []Track{track},
		Links:  included,
	})
}

func (f *TracksFeature) TrackAudio(w http.ResponseWriter, r *http.Request) {
}

func (f *TracksFeature) TrackImages(w http.ResponseWriter, r *http.Request) {
}

func (f *TracksFeature) TrackImage(w http.ResponseWriter, r *http.Request) {
}
