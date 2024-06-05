package api

import (
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sdgondola/wasaphoto/service/database"
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
	follower := ps.ByName("userID")
	if follower == "" {
		http.Error(w, "Bad request: no userID provided", http.StatusBadRequest)
		return
	} else if follower != token {
		http.Error(w, "Bad request: bad userID", http.StatusBadRequest)
		return
	}
	toUnfollow := ps.ByName("userID")
	if toUnfollow == "" {
		http.Error(w, "Bad request: no userID provided", http.StatusBadRequest)
		return
	}
	if toUnfollow == token {
		http.Error(w, "Bad request: you cannot unfollow yourself!", http.StatusBadRequest)
		return
	}
	err = rt.db.Unfollow(token, toUnfollow)
	if errors.Is(err, database.ErrUserNotFound) {
		http.Error(w, "Bad request: no such user", http.StatusBadRequest)
	} else if errors.Is(err, database.ErrUserIsBlocked) {
		http.Error(w, "Forbidden: user blocked you!", http.StatusForbidden)
	} else if errors.Is(err, database.ErrNotFollowing) {
		http.Error(w, "Bad request: not following", http.StatusBadRequest)
	} else if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
