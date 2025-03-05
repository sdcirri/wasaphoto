package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/sdcirri/wasaphoto/service/database"
)

func (rt *_router) unlikeComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	userID, err := strconv.ParseInt(ps.ByName("userID"), 10, 64)
	if err != nil {
		http.Error(w, "Bad userID", http.StatusBadRequest)
		return
	}
	if token != userID {
		http.Error(w, "Error: trying to unlike as somebody else", http.StatusForbidden)
		return
	}
	commentID, err := strconv.ParseInt(ps.ByName("commentID"), 10, 64)
	if err != nil {
		http.Error(w, "Bad commentID", http.StatusBadRequest)
		return
	}

	err = rt.db.UnlikeComment(token, commentID)
	if errors.Is(err, database.ErrUserNotFound) {
		http.Error(w, "Bad authentication token", http.StatusBadRequest)
	} else if errors.Is(err, database.ErrCommentNotFound) {
		http.Error(w, database.ErrCommentNotFound.Error(), http.StatusNotFound)
	} else if errors.Is(err, database.ErrUserIsBlocked) {
		http.Error(w, "Cannot like comment: user blocked you!", http.StatusForbidden)
	} else if err != nil && !errors.Is(err, database.ErrAlreadyLiked) { // We can safely ignore that as it's likely some duplicate request
		rt.internalServerError(err, w)
	} else {
		c, err := rt.db.GetComment(commentID)
		if err != nil {
			rt.internalServerError(err, w)
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("content-type", "text/plain")
		_, err = w.Write([]byte(strconv.FormatInt(int64(c.Likes), 10)))
		if err != nil {
			rt.internalServerError(err, w)
		}
	}
}
