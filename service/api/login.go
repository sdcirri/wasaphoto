package api

import (
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sdgondola/wasaphoto/service/database"
)

func (rt *_router) login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// First check if the user is already logged in
	username, err := rt.getAuthToken(r)
	if errors.Is(err, ErrNoAuth) {
		http.Error(w, "Bad request: no auth token provided", http.StatusBadRequest)
		return
	} else if errors.Is(err, database.ErrUserNotFound) {
		if len(username) < 3 || len(username) > 127 {
			http.Error(w, "Invalid username, usernames should be between 3 and 127 characters long", http.StatusBadRequest)
			return
		}
		err = rt.db.RegisterUser(username)
		if err != nil {
			rt.internalServerError(err, w)
			return
		}
		w.WriteHeader(http.StatusCreated)
	} else if err != nil {
		rt.internalServerError(err, w)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "WASASESSIONID",
		Value: username,
		Path:  "/",
	})
	w.Header().Set("content-type", "text-plain")
	_, err = w.Write([]byte(username))
	if err != nil {
		rt.internalServerError(err, w)
		return
	}
}
