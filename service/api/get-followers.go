package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sdgondola/wasaphoto/service/database"
)

func (rt *_router) getFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	token, err := rt.getAuthToken(r)
	if errors.Is(err, ErrNoAuth) {
		http.Error(w, "Unauthenticated", http.StatusUnauthorized)
		return
	} else if errors.Is(err, database.ErrUserNotFound) {
		http.Error(w, "Bad authentication token", http.StatusBadRequest)
		return
	} else if err != nil {
		rt.internalServerError(err, w)
		return
	} else if token != ps.ByName("userID") {
		http.Error(w, "Forbidden: cannot view somebody else's followers", http.StatusForbidden)
		return
	}

	followers, err := rt.db.GetFollowers(token)
	if errors.Is(err, database.ErrUserNotFound) {
		// Very unlikely but still possible
		http.Error(w, "Bad authentication token", http.StatusBadRequest)
	} else if err != nil {
		rt.internalServerError(err, w)
	} else {
		j, err := json.Marshal(followers)
		if err != nil {
			rt.internalServerError(err, w)
		} else {
			w.Header().Set("content-type", "application/json")
			_, err = w.Write(j)
			if err != nil {
				rt.internalServerError(err, w)
			}
		}
	}
}
