package aura

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

type Feature struct {
}

// get a URL variable
func (f *Feature) getVar(r *http.Request, id string) string {
	return mux.Vars(r)[id]
}

// get query params
func (f *Feature) getQueryParams(r *http.Request) map[string]string {
	params := make(map[string]string)
	for key, values := range r.URL.Query() {
		params[key] = values[0]
	}
	return params
}

// get include from params
func (f *Feature) getInclude(params *map[string]string) map[string]bool {
	include := map[string]bool{}
	if pInclude, ok := (*params)["include"]; ok == true {
		for _, key := range strings.Split(pInclude, ",") {
			include[key] = true
		}
		delete(*params, "include")
	}
	return include
}

// get limit
func (f *Feature) getLimit(params *map[string]string) int {
	limit := -1
	if pLimit, ok := (*params)["limit"]; ok == true {
		var err error
		limit, err = strconv.Atoi(pLimit)
		if err != nil {
			limit = -1
		}
		delete(*params, "limit")
	}
	return limit
}

// printJson
func (f *Feature) sendJSON(w http.ResponseWriter, obj interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return json.NewEncoder(w).Encode(obj)
}
