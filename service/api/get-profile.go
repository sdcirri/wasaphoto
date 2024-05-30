package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sdgondola/wasaphoto/service/database"
)

func (rt *_router) getProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	canViewPosts := true
	token, err := rt.getAuthToken(r)
	if errors.Is(err, ErrNoAuth) {
		// Let unauthenticated users preview the profile
		canViewPosts = false
	} else if errors.Is(err, database.ErrUserNotFound) {
		http.Error(w, "Bad auth token", http.StatusBadRequest)
		return
	} else if err != nil {
		rt.internalServerError(err, w)
		return
	}

	toView := ps.ByName("username")
	profile, err := rt.db.GetAccount(token, toView)
	if errors.Is(err, database.ErrUserIsBlocked) {
		http.Error(w, "Forbidden: user blocked you!", http.StatusForbidden)
	} else if errors.Is(err, database.ErrUserNotFound) {
		http.Error(w, "User not found", http.StatusNotFound)
	} else if err != nil {
		rt.internalServerError(err, w)
	} else {
		if !canViewPosts {
			profile.Posts = make([]int64, 0)
		}
		j, err := json.Marshal(profile)
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
