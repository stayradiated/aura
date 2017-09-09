package aura

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func GetRoutes(h Features) Routes {
	routes := make(Routes, 0)

	if h.Tracks != nil {
		routes = append(routes, h.Tracks.Routes()...)
	}

	if h.Artists != nil {
		routes = append(routes, h.Artists.Routes()...)
	}

	if h.Albums != nil {
		routes = append(routes, h.Albums.Routes()...)
	}

	if h.Playlists != nil {
		routes = append(routes, h.Playlists.Routes()...)
	}

	return routes
}
