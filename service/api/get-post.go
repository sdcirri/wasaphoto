package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/sdgondola/wasaphoto/service/database"
)

// Here we don't want to return the likes explicitly, but rather the like count
type PostFE struct {
	PostID    int64   `json:"postID"`
	ImageB64  string  `json:"imageB64"`
	PubTime   string  `json:"pubTime"`
	Caption   string  `json:"caption"`
	Author    string  `json:"author"`
	LikeCount int     `json:"likeCount"`
	Comments  []int64 `json:"comments"`
}

func post2postFE(post database.Post) PostFE {
	return PostFE{
		PostID:    post.PostID,
		ImageB64:  post.ImageB64,
		PubTime:   post.PubTime,
		Caption:   post.Caption,
		Author:    post.Author,
		LikeCount: len(post.Likes),
		Comments:  post.Comments,
	}
}

func (rt *_router) getPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	post, err := rt.db.GetPost(token, postID)
	if err == database.ErrUserNotFound {
		http.Error(w, "Bad authentication token", http.StatusBadRequest)
	} else if errors.Is(err, database.ErrPostNotFound) {
		http.Error(w, database.ErrPostNotFound.Error(), http.StatusNotFound)
	} else if errors.Is(err, database.ErrUserIsBlocked) {
		http.Error(w, "Cannot view post: user blocked you!", http.StatusForbidden)
	} else if err != nil {
		rt.internalServerError(err, w)
	} else {
		j, err := json.Marshal(post2postFE(post))
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
