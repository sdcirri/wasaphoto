package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/sdcirri/wasaphoto/service/database"
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
	userID, err := strconv.ParseInt(ps.ByName("userID"), 10, 64)
	if err != nil {
		http.Error(w, "Bad userID", http.StatusBadRequest)
		return
	}
	if token != userID {
		http.Error(w, "Forbidden: cannot delete somebosy else's followers", http.StatusForbidden)
		return
	}
	toRm, err := strconv.ParseInt(ps.ByName("toRemoveID"), 10, 64)
	if err != nil {
		http.Error(w, "Bad userID", http.StatusBadRequest)
		return
	}
	if toRm == token {
		http.Error(w, "Bad request: you cannot follow yourself!", http.StatusBadRequest)
		return
	}
	err = rt.db.RmFollower(token, toRm)
	if errors.Is(err, database.ErrUserNotFound) {
		http.Error(w, "Error: no such user", http.StatusNotFound)
	} else if errors.Is(err, database.ErrNotFollowing) {
		http.Error(w, "Error: user does not follow you!", http.StatusNotFound)
	} else if errors.Is(err, database.ErrAlreadyFollowing) {
		http.Error(w, "Bad request: already following", http.StatusBadRequest)
	} else if err != nil {
		rt.internalServerError(err, w)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
