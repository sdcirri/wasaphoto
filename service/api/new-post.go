package api

import (
	"encoding/json"
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
	idc, err := r.Cookie("WASASESSIONID")
	if err == http.ErrNoCookie {
		http.Error(w, "Unauthenticated", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	var postParams PostParams
	id := idc.Value
	err = json.NewDecoder(r.Body).Decode(&postParams)
	if err != nil {
		http.Error(w, "Bad request: malformed json: "+err.Error(), http.StatusBadRequest)
		return
	}
	if len(postParams.Caption) > 2048 {
		http.Error(w, "Bad request: caption too long", http.StatusBadRequest)
		return
	}

	postID, err := rt.db.NewPost(id, postParams.Image, postParams.Caption)
	if err == database.ErrUserNotFound {
		http.Error(w, "Bad request: hacking attempt?!", http.StatusBadRequest)
	}
	if err == database.ErrBadImage {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "text/plain")
	_, _ = w.Write([]byte(strconv.FormatInt(postID, 10)))
}
