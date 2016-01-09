package aura

import (
	"encoding/json"
	"net/http"
)

type Feature struct {
}

// printJson
func (f *Feature) sendJSON(w http.ResponseWriter, obj interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return json.NewEncoder(w).Encode(obj)
}
