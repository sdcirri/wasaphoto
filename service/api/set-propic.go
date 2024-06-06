package api

import (
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/sdgondola/wasaphoto/service/database"
)

func (rt *_router) setProPic(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	token, err := rt.getAuthToken(r)
	if errors.Is(err, database.ErrUserNotFound) {
		http.Error(w, "Error: bad authentication token", http.StatusBadRequest)
		return
	} else if errors.Is(err, ErrNoAuth) {
		http.Error(w, "Error: no auth token provided", http.StatusUnauthorized)
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
	if userID != token {
		http.Error(w, "Error: you cannot set somebody else's profile picture", http.StatusForbidden)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		rt.internalServerError(err, w)
		return
	}
	err = rt.db.SetProPic(userID, string(body[:]))
	if errors.Is(err, database.ErrUserNotFound) {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if errors.Is(err, database.ErrBadImage) {
		http.Error(w, "Bad image", http.StatusBadRequest)
		return
	} else if err != nil {
		rt.internalServerError(err, w)
		return
	}

	w.Header().Set("content-type", "text/plain")
	_, err = w.Write([]byte(ps.ByName("userID"))) // We spare to use strconv.FormatInt()
	if err != nil {
		rt.internalServerError(err, w)
	}
}
