package aura

import (
	"github.com/gorilla/mux"
)

func NewRouter(features Features) *mux.Router {
	router := mux.NewRouter()

	// load routes
	for _, route := range GetRoutes(features) {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
