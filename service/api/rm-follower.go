package api

import (
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sdgondola/wasaphoto/service/database"
)

func (rt *_router) rmFollower(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	}
	toRm := ps.ByName("username")
	if toRm == "" {
		http.Error(w, "Bad request: no username provided", http.StatusBadRequest)
		return
	}
	if toRm == token {
		http.Error(w, "Bad request: you cannot follow yourself!", http.StatusBadRequest)
		return
	}
	err = rt.db.RmFollower(token, toRm)
	if err == database.ErrUserNotFound {
		http.Error(w, "Error: no such user", http.StatusNotFound)
	} else if err == database.ErrNotFollowing {
		http.Error(w, "Error: user does not follow you!", http.StatusNotFound)
	} else if err == database.ErrAlreadyFollowing {
		http.Error(w, "Bad request: already following", http.StatusBadRequest)
	} else if err != nil {
		rt.internalServerError(err, w)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
