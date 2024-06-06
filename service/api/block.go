package api

import (
	"errors"
	"net/http"
	"strconv"

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
	blocker, err := strconv.ParseInt(ps.ByName("userID"), 10, 64)
	if err != nil || blocker != token {
		http.Error(w, "Bad userID", http.StatusBadRequest)
		return
	}

	toBlock, err := strconv.ParseInt(ps.ByName("toBlockID"), 10, 64)
	if err != nil {
		http.Error(w, "Baad userID", http.StatusBadRequest)
		return
	}
	if toBlock == token {
		http.Error(w, "Bad request: you cannot block yourself!", http.StatusBadRequest)
		return
	}
	err = rt.db.Block(token, toBlock)
	if errors.Is(err, database.ErrUserNotFound) {
		http.Error(w, "Error: user not found", http.StatusNotFound)
		return
	} else if errors.Is(err, database.ErrAlreadyBlocked) {
		http.Error(w, "Bad request: user already blocked", http.StatusBadRequest)
		return
	} else if err != nil {
		rt.internalServerError(err, w)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "text/plain")
	_, err = w.Write([]byte(strconv.FormatInt(toBlock, 10)))
	if err != nil {
		rt.internalServerError(err, w)
	}
}
