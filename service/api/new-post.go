package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/sdgondola/wasaphoto/service/database"
)

type PostParams struct {
	Image   string `json:"image"`
	Caption string `json:"caption"`
}

func (rt *_router) newPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
		http.Error(w, "Forbidden: cannot post as somebody else", http.StatusForbidden)
		return
	}

	var postParams PostParams
	err = json.NewDecoder(r.Body).Decode(&postParams)
	if err != nil {
		http.Error(w, "Bad request: malformed json: "+err.Error(), http.StatusBadRequest)
		return
	}
	if len(postParams.Image) > 3000000 {
		http.Error(w, "bad request: image too big, images up to 2 MB are supported", http.StatusBadRequest)
		return
	}
	if len(postParams.Caption) > 2048 {
		http.Error(w, "Bad request: caption too long", http.StatusBadRequest)
		return
	}

	postID, err := rt.db.NewPost(token, postParams.Image, postParams.Caption)
	if errors.Is(err, database.ErrUserNotFound) {
		http.Error(w, "Bad authentication token", http.StatusBadRequest)
	}
	if errors.Is(err, database.ErrBadImage) {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if err != nil {
		rt.internalServerError(err, w)
		return
	}

	w.Header().Set("content-type", "text/plain")
	_, err = w.Write([]byte(strconv.FormatInt(postID, 10)))
	if err != nil {
		rt.internalServerError(err, w)
	}
}
