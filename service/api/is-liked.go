package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/sdcirri/wasaphoto/service/database"
)

func (rt *_router) isLiked(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	token, err := rt.getAuthToken(r)
	if errors.Is(err, ErrNoAuth) {
		http.Error(w, "Unauthenticated", http.StatusUnauthorized)
		return
	} else if errors.Is(err, database.ErrUserNotFound) {
		http.Error(w, "Bad authentication token", http.StatusUnauthorized)
		return
	} else if err != nil {
		rt.internalServerError(err, w)
		return
	}
	userID, err := strconv.ParseInt(ps.ByName("userID"), 10, 64)
	if err != nil || userID != token {
		http.Error(w, "Bad authentication token", http.StatusUnauthorized)
		return
	}
	postID, err := strconv.ParseInt(ps.ByName("postID"), 10, 64)
	if err != nil {
		http.Error(w, database.ErrPostNotFound.Error(), http.StatusNotFound)
		return
	}
	liked, err := rt.db.IsLiked(userID, postID)
	if err != nil {
		rt.internalServerError(err, w)
	} else {
		var body []byte
		if liked {
			body = []byte("true")
		} else {
			body = []byte("false")
		}
		w.Header().Set("content-type", "text/plain")
		_, err = w.Write(body)
		if err != nil {
			rt.internalServerError(err, w)
		}
	}
}
