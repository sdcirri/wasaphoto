package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) SearchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	q := r.URL.Query().Get("q")
	if q == "" {
		http.Error(w, "Bad request: no search query provided", http.StatusBadRequest)
		return
	}
	res, err := rt.db.SearchUser(q)
	if err != nil {
		rt.internalServerError(err, w)
		return
	}
	j, err := json.Marshal(res)
	if err != nil {
		rt.internalServerError(err, w)
		return
	}
	w.Header().Set("content-type", "application/json")
	_, err = w.Write(j)
	if err != nil {
		rt.internalServerError(err, w)
	}
}
