package api

import (
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sdgondola/wasaphoto/service/database"
)

func (rt *_router) block(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	toBlock := ps.ByName("username")
	if toBlock == "" {
		http.Error(w, "Bad request: no username provided", http.StatusBadRequest)
		return
	}
	if toBlock == token {
		http.Error(w, "Bad request: you cannot block yourself!", http.StatusBadRequest)
		return
	}
	err = rt.db.Block(token, toBlock)
	if errors.Is(err, database.ErrUserNotFound) {
		http.Error(w, "Error: user not found", http.StatusNotFound)
	} else if errors.Is(err, database.ErrAlreadyBlocked) {
		http.Error(w, "Bad request: user already blocked", http.StatusBadRequest)
	} else if err != nil {
		rt.internalServerError(err, w)
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("content-type", "text/plain")
		_, err = w.Write([]byte(toBlock))
		if err != nil {
			rt.internalServerError(err, w)
		}
	}
}
