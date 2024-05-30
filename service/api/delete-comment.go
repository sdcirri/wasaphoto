package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/sdgondola/wasaphoto/service/database"
)

func (rt *_router) deleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	if token != ps.ByName("userID") {
		http.Error(w, "Error: trying to delete somebody else's comment", http.StatusForbidden)
		return
	}
	commentID, err := strconv.ParseInt(ps.ByName("commentID"), 10, 64)
	if err != nil {
		http.Error(w, "Bad comment ID: "+ps.ByName("commentID"), http.StatusBadRequest)
		return
	}

	err = rt.db.DeleteComment(token, commentID)
	if errors.Is(err, database.ErrUserNotFound) {
		http.Error(w, "Bad request: no such user", http.StatusBadRequest)
	} else if errors.Is(err, database.ErrUserIsNotAuthor) {
		http.Error(w, "Error: trying to delete somebody else's comment", http.StatusForbidden)
	} else if errors.Is(err, database.ErrCommentNotFound) {
		http.Error(w, database.ErrCommentNotFound.Error(), http.StatusNotFound)
	} else if err != nil {
		rt.internalServerError(err, w)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
