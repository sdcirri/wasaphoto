package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/sdcirri/wasaphoto/service/database"
)

func (rt *_router) unfollow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	follower, err := strconv.ParseInt(ps.ByName("userID"), 10, 64)
	if err != nil || follower != token {
		http.Error(w, "Bad userID", http.StatusBadRequest)
		return
	}
	toUnfollow, err := strconv.ParseInt(ps.ByName("toUnfollowID"), 10, 64)
	if err != nil {
		http.Error(w, "Bad userID", http.StatusBadRequest)
		return
	}
	if toUnfollow == token {
		http.Error(w, "Bad request: you cannot unfollow yourself!", http.StatusBadRequest)
		return
	}
	err = rt.db.Unfollow(token, toUnfollow)
	if errors.Is(err, database.ErrUserNotFound) {
		http.Error(w, database.ErrUserNotFound.Error(), http.StatusNotFound)
	} else if errors.Is(err, database.ErrUserIsBlocked) {
		http.Error(w, "Forbidden: user blocked you!", http.StatusForbidden)
	} else if errors.Is(err, database.ErrNotFollowing) {
		http.Error(w, database.ErrNotFollowing.Error(), http.StatusNotFound)
	} else if err != nil {
		rt.internalServerError(err, w)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
