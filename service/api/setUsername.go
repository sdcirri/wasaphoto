package api

import (
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sdgondola/wasaphoto/service/database"
)

func (rt *_router) setUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Bad request: no username provided", http.StatusBadGateway)
		return
	}
	taken, err := rt.db.UsernameTaken(username)
	if err != nil {
		rt.internalServerError(err, w)
		return
	}
	if taken {
		http.Error(w, "Username already taken", http.StatusForbidden)
		return
	}

	err = rt.db.SetUsername(token, username)
	if err != nil {
		rt.internalServerError(err, w)
		return
	}

	w.Header().Set("content-type", "text/plain")
	_, err = w.Write([]byte(username))
	if err != nil {
		rt.internalServerError(err, w)
	}
}
