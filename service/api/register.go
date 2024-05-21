package api

import (
	"github.com/julienschmidt/httprouter"
    "github.com/sdgondola/wasaphoto/service/database"
	"net/http"
)

func (rt *_router) register(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Bad request: username must be provided", http.StatusBadRequest)
		return
	}
	exists, err := rt.db.UserExists(username)
	if err != nil {
		http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
		return
	}
    if exists {
        http.Error(w, "Username already taken", http.StatusForbidden)
        return
    }
    err = rt.db.RegisterUser(username)
    if err == database.ErrUserAlreadyExists {
        http.Error(w, "Username already taken", http.StatusForbidden)
        return
    }
    if err != nil {
		http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
		return
	}
    http.SetCookie(w, &http.Cookie{
		Name: "WASASESSIONID",
		Value: username,
		Path: "/",
	})
	w.WriteHeader(http.StatusCreated)
}
