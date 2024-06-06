package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/sdgondola/wasaphoto/service/database"
)

func (rt *_router) unblock(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	blocker, err := strconv.ParseInt(ps.ByName("userID"), 64, 10)
	if err != nil || blocker != token {
		http.Error(w, "Bad userID", http.StatusBadRequest)
		return
	}
	toUnblock, err := strconv.ParseInt(ps.ByName("toUnblockID"), 64, 10)
	if err != nil {
		http.Error(w, "Bad userID", http.StatusBadRequest)
		return
	}
	if toUnblock == token {
		http.Error(w, "Bad request: you cannot block yourself!", http.StatusBadRequest)
		return
	}
	err = rt.db.Unblock(token, toUnblock)
	if errors.Is(err, database.ErrUserNotFound) {
		http.Error(w, database.ErrUserNotFound.Error(), http.StatusNotFound)
	} else if errors.Is(err, database.ErrUserNotBlocked) {
		http.Error(w, database.ErrUserNotBlocked.Error(), http.StatusNotFound)
	} else if err != nil {
		rt.internalServerError(err, w)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
