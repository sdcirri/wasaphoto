package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/sdgondola/wasaphoto/service/database"
)

func (rt *_router) getComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	commentID, err := strconv.ParseInt(ps.ByName("commentID"), 10, 64)
	if err != nil {
		http.Error(w, database.ErrPostNotFound.Error(), http.StatusNotFound)
		return
	}
	comment, err := rt.db.GetComment(commentID)
	if errors.Is(err, database.ErrCommentNotFound) {
		http.Error(w, database.ErrCommentNotFound.Error(), http.StatusNotFound)
	} else if err != nil {
		rt.internalServerError(err, w)
	} else {
		blocked, err := rt.db.IsBlockedBy(token, comment.Author)
		if err != nil {
			rt.internalServerError(err, w)
			return
		}
		if blocked {
			http.Error(w, "Cannot view comment: user blocked you!", http.StatusForbidden)
			return
		}
		j, err := json.Marshal(comment)
		if err != nil {
			rt.internalServerError(err, w)
			return
		}
		w.Header().Set("content-type", "application/json")
		_, err = w.Write(j)
		if err != nil {
			rt.internalServerError(err, w)
		}
	}
}
