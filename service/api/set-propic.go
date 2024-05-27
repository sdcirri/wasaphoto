package api

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sdgondola/wasaphoto/service/database"
)

func (rt *_router) setProPic(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := r.Cookie("WASASESSIONID")
	if err == http.ErrNoCookie {
		http.Error(w, "Unauthenticated", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = rt.db.SetProPic(id.Value, string(body[:]))
	if err == database.ErrUserNotFound {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err == database.ErrBadImage {
		http.Error(w, "Bad image", http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
