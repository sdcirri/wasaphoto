package api

import (
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/sdgondola/wasaphoto/service/database"
)

func (rt *_router) commentPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	if err != nil {
		http.Error(w, "Bad userID", http.StatusUnauthorized)
		return
	}
	if token != userID {
		http.Error(w, "Error: trying to comment as somebody else", http.StatusForbidden)
		return
	}
	postID, err := strconv.ParseInt(ps.ByName("postID"), 10, 64)
	if err != nil {
		http.Error(w, "Bad postID", http.StatusNotFound)
		return
	}
	textRaw, err := io.ReadAll(r.Body)
	if err != nil {
		rt.internalServerError(err, w)
		return
	}
	text := string(textRaw[:])
	if text == "" {
		http.Error(w, "Bad request: empty comment", http.StatusBadRequest)
		return
	}
	if len(text) > 2048 {
		http.Error(w, "Bad request: comment too long", http.StatusBadRequest)
		return
	}

	cID, err := rt.db.CommentPost(token, postID, text)
	if errors.Is(err, database.ErrUserNotFound) {
		http.Error(w, "Bad authentication token", http.StatusBadRequest)
	} else if errors.Is(err, database.ErrPostNotFound) {
		http.Error(w, database.ErrPostNotFound.Error(), http.StatusNotFound)
	} else if errors.Is(err, database.ErrUserIsBlocked) {
		http.Error(w, "Cannot like post: user blocked you!", http.StatusForbidden)
	} else if err != nil {
		rt.internalServerError(err, w)
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("content-type", "text/plain")
		_, err = w.Write([]byte(strconv.FormatInt(cID, 10)))
		if err != nil {
			rt.internalServerError(err, w)
		}
	}
}
