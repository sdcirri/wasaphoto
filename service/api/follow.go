package api

import (
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sdgondola/wasaphoto/service/database"
)

func (rt *_router) follow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	toFollow := ps.ByName("userID")
	if toFollow == "" {
		http.Error(w, "Bad request: no userID provided", http.StatusBadRequest)
		return
	}
	if toFollow == token {
		http.Error(w, "Bad request: you cannot follow yourself!", http.StatusBadRequest)
		return
	}
	err = rt.db.Follow(token, toFollow)
	if errors.Is(err, database.ErrUserNotFound) {
		http.Error(w, "Bad request: no such user", http.StatusBadRequest)
	} else if errors.Is(err, database.ErrUserIsBlocked) {
		http.Error(w, "Forbidden: user blocked you!", http.StatusForbidden)
	} else if errors.Is(err, database.ErrAlreadyFollowing) {
		http.Error(w, "Bad request: already following", http.StatusBadRequest)
	} else if err != nil {
		rt.internalServerError(err, w)
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("content-type", "text/plain")
		_, err = w.Write([]byte(toFollow))
		if err != nil {
			rt.internalServerError(err, w)
		}
	}
}
