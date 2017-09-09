package aura

import (
	"io"
	"log"
	"net/http"
)

type PlaylistsInterface interface {
	FilterPlaylists(params map[string]string, include Include, limit int, cont string) (Entities, error)
	PlaylistWithID(playlistID string, include Include) (Entities, error)
	PlaylistImageWithID(playlistID, imageID string, w io.Writer) error
}

type PlaylistsFeature struct {
	Feature
	PlaylistsInterface
}

func (f *PlaylistsFeature) Routes() Routes {
	return Routes{
		Route{
			"AllPlaylists",
			"GET", "/playlists", f.getPlaylists,
		},
		Route{
			"Playlist",
			"GET", "/playlists/{playlistID}", f.getPlaylistWithID,
		},
		Route{
			"PlaylistImages",
			"GET", "/playlists/{playlistID}/images", f.getPlaylistImage,
		},
		Route{
			"PlaylistImage",
			"GET", "/playlists/{playlistID}/images/{imageID}", f.getPlaylistImage,
		},
	}
}

func (f *PlaylistsFeature) getPlaylists(w http.ResponseWriter, r *http.Request) {
	params := f.getQueryParams(r)
	include := f.getInclude(&params)
	limit := f.getLimit(&params)

	entities, err := f.FilterPlaylists(params, include, limit, "")
	if err != nil {
		log.Fatal(err)
	}

	f.sendJSON(w, Response{
		Result:   entities.Playlists.IDs(),
		Entities: entities.asMap(),
	})
}

func (f *PlaylistsFeature) getPlaylistWithID(w http.ResponseWriter, r *http.Request) {
	playlistID := f.getVar(r, "playlistID")
	params := f.getQueryParams(r)
	include := f.getInclude(&params)

	entities, err := f.PlaylistWithID(playlistID, include)
	if err != nil {
		log.Fatal(err)
	}

	f.sendJSON(w, Response{
		Result:   entities.Playlists.IDs(),
		Entities: entities.asMap(),
	})
}

func (f *PlaylistsFeature) getPlaylistImages(w http.ResponseWriter, r *http.Request) {
}

func (f *PlaylistsFeature) getPlaylistImage(w http.ResponseWriter, r *http.Request) {
	playlistID := f.getVar(r, "playlistID")
	imageID := f.getVar(r, "imageID")
	f.PlaylistImageWithID(playlistID, imageID, w)
}
