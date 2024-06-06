package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/sdgondola/wasaphoto/service/database"
)

func (rt *_router) likePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
		http.Error(w, "Error: trying to like as somebody else", http.StatusForbidden)
		return
	}
	postID, err := strconv.ParseInt(ps.ByName("postID"), 10, 64)
	if err != nil {
		http.Error(w, "Bad postID", http.StatusBadRequest)
		return
	}

	err = rt.db.LikePost(token, postID)
	if errors.Is(err, database.ErrUserNotFound) {
		http.Error(w, "Bad authentication token", http.StatusBadRequest)
	} else if errors.Is(err, database.ErrPostNotFound) {
		http.Error(w, database.ErrPostNotFound.Error(), http.StatusNotFound)
	} else if errors.Is(err, database.ErrUserIsBlocked) {
		http.Error(w, "Cannot like post: user blocked you!", http.StatusForbidden)
	} else if err != nil && !errors.Is(err, database.ErrAlreadyLiked) { // We can safely ignore that as it's likely some duplicate request
		rt.internalServerError(err, w)
	} else {
		post, err := rt.db.GetPost(token, postID)
		if err != nil {
			rt.internalServerError(err, w)
		}
		likeCount := len(post.Likes)
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("content-type", "text-plain")
		_, err = w.Write([]byte(strconv.FormatInt(int64(likeCount), 10)))
		if err != nil {
			rt.internalServerError(err, w)
		}
	}
}
