package aura

import (
	"log"
	"net/http"
)

type TracksInterface interface {
	FilterTracks(params map[string]string, include Include, limit int, cont string) (Entities, error)
	TrackWithID(trackID string, include Include) (Entities, error)
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

func (f *TracksFeature) getTracks(w http.ResponseWriter, r *http.Request) {
	params := f.getQueryParams(r)
	include := f.getInclude(&params)
	limit := f.getLimit(&params)

	entities, err := f.FilterTracks(params, include, limit, "")
	if err != nil {
		log.Fatal(err)
	}

	f.sendJSON(w, Response{
		Result:   entities.Tracks.IDs(),
		Entities: entities.asMap(),
	})
}

func (f *TracksFeature) getTrackWithID(w http.ResponseWriter, r *http.Request) {
	trackID := f.getVar(r, "trackID")
	params := f.getQueryParams(r)
	include := f.getInclude(&params)

	entities, err := f.TrackWithID(trackID, include)
	if err != nil {
		log.Fatal(err)
	}

	f.sendJSON(w, Response{
		Result:   entities.Tracks.IDs(),
		Entities: entities.asMap(),
	})
}

func (f *TracksFeature) TrackAudio(w http.ResponseWriter, r *http.Request) {
}

func (f *TracksFeature) TrackImages(w http.ResponseWriter, r *http.Request) {
}

func (f *TracksFeature) TrackImage(w http.ResponseWriter, r *http.Request) {
}
