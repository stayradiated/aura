package aura

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type AlbumsInterface interface {
	GetAllAlbums(limit int, cont string) ([]Album, error)
	GetAlbumByID(albumID string) (Album, error)
}

type AlbumsFeature struct {
	Feature
	AlbumsInterface
}

func (f *AlbumsFeature) Routes() Routes {
	return Routes{
		Route{
			"AllAlbums",
			"GET", "/albums", f.AllAlbums,
		},
		Route{
			"Album",
			"GET", "/albums/{albumID}", f.Album,
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

func (f *AlbumsFeature) AllAlbums(w http.ResponseWriter, r *http.Request) {
	albums, err := f.GetAllAlbums(-1, "")
	if err != nil {
		log.Fatal(err)
	}

	f.sendJSON(w, albums)
}

func (f *AlbumsFeature) Album(w http.ResponseWriter, r *http.Request) {
	albumID := mux.Vars(r)["albumID"]

	album, err := f.GetAlbumByID(albumID)
	if err != nil {
		log.Fatal(err)
	}

	f.sendJSON(w, album)
}

func (f *AlbumsFeature) AlbumImages(w http.ResponseWriter, r *http.Request) {
}

func (f *AlbumsFeature) AlbumImage(w http.ResponseWriter, r *http.Request) {
}
