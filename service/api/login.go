package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// First check if the user is already logged in
	id, err := r.Cookie("WASASESSIONID")
	if err == http.ErrNoCookie {
		// If not, let's check the ID in the query
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
		if !exists {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name: "WASASESSIONID",
			Value: username,
			Path: "/",
		})
		w.WriteHeader(http.StatusOK)
		return
	} else if err != nil {
    	http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
    	return
	}

	// Check if cookie is not valid
	exists, err := rt.db.UserExists(id.Value)
	if err != nil {
    	http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
    	return
	}
	if !exists {
		// This is kinda suspicious, likely a forged cookie
		http.Error(w, "Bad request: hacking attempt?!", http.StatusBadRequest)
		return
	}

	// Cookie is already present and valid, we don't need to do anything
}
