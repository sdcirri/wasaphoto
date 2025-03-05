package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sdcirri/wasaphoto/service/database"
)

func (rt *_router) getFeed(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	token, err := rt.getAuthToken(r)
	if errors.Is(err, ErrNoAuth) {
		http.Error(w, "Unauthenticated", http.StatusUnauthorized)
		return
	} else if errors.Is(err, database.ErrUserNotFound) {
		http.Error(w, "Bad authentication token", http.StatusUnauthorized)
		return
	} else if err != nil {
		rt.internalServerError(err, w)
		return
	}
	feed, err := rt.db.GetFeed(token)
	if errors.Is(err, database.ErrUserNotFound) {
		http.Error(w, "Bad authentication token", http.StatusUnauthorized)
	} else if err != nil {
		rt.internalServerError(err, w)
	} else {
		j, err := json.Marshal(feed)
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
}
