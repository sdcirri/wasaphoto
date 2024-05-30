package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/sdgondola/wasaphoto/service/database"
)

func (rt *_router) rmPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	postID, err := strconv.ParseInt(ps.ByName("postID"), 10, 64)
	if err != nil {
		http.Error(w, database.ErrPostNotFound.Error(), http.StatusNotFound)
		return
	}

	err = rt.db.RmPost(token, postID)
	if errors.Is(err, database.ErrPostNotFound) {
		http.Error(w, database.ErrPostNotFound.Error(), http.StatusNotFound)
	} else if errors.Is(err, database.ErrUserIsNotAuthor) {
		http.Error(w, database.ErrUserIsNotAuthor.Error(), http.StatusForbidden)
	} else if err != nil {
		rt.internalServerError(err, w)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
