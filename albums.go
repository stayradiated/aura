package aura

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type AlbumsInterface interface {
	FilterAlbums(params map[string]string, include Include, limit int, cont string) (Entities, error)
	AlbumWithID(albumID string, include Include) (Entities, error)
	AlbumImageWithID(albumID, imageID string, w io.Writer) error
}

type AlbumsFeature struct {
	Feature
	AlbumsInterface
}

func (f *AlbumsFeature) Routes() Routes {
	return Routes{
		Route{
			"AllAlbums",
			"GET", "/albums", f.getAlbums,
		},
		Route{
			"Album",
			"GET", "/albums/{albumID}", f.getAlbumWithID,
		},
		Route{
			"AlbumImages",
			"GET", "/albums/{albumID}/images", f.AlbumImages,
		},
		Route{
			"AlbumImage",
			"GET", "/albums/{albumID}/images/{imageID}", f.AlbumImage,
		},
	}
}

func (f *AlbumsFeature) getAlbums(w http.ResponseWriter, r *http.Request) {
	params := f.getQueryParams(r)
	include := f.getInclude(&params)
	limit := f.getLimit(&params)

	entities, err := f.FilterAlbums(params, include, limit, "")
	if err != nil {
		log.Fatal(err)
	}

	f.sendJSON(w, Response{
		Result:   entities.Albums.IDs(),
		Entities: entities.asMap(),
	})
}

func (f *AlbumsFeature) getAlbumWithID(w http.ResponseWriter, r *http.Request) {
	albumID := mux.Vars(r)["albumID"]
	params := f.getQueryParams(r)
	include := f.getInclude(&params)

	entities, err := f.AlbumWithID(albumID, include)
	if err != nil {
		log.Fatal(err)
	}

	f.sendJSON(w, Response{
		Result:   entities.Albums.IDs(),
		Entities: entities.asMap(),
	})
}

func (f *AlbumsFeature) AlbumImages(w http.ResponseWriter, r *http.Request) {
}

func (f *AlbumsFeature) AlbumImage(w http.ResponseWriter, r *http.Request) {
	albumID := mux.Vars(r)["albumID"]
	imageID := mux.Vars(r)["imageID"]
	f.AlbumImageWithID(albumID, imageID, w)
}
