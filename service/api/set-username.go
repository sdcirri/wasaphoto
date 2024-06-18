package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/sdgondola/wasaphoto/service/database"
)

func (rt *_router) setUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	token, err := rt.getAuthToken(r)
	if errors.Is(err, database.ErrUserNotFound) {
		http.Error(w, "Error: bad authentication token", http.StatusUnauthorized)
		return
	} else if errors.Is(err, ErrNoAuth) {
		http.Error(w, "Error: no auth token provided", http.StatusUnauthorized)
		return
	} else if err != nil {
		rt.internalServerError(err, w)
		return
	}
	userID, err := strconv.ParseInt(ps.ByName("userID"), 10, 64)
	if err != nil || userID != token {
		http.Error(w, "Bad userID", http.StatusBadRequest)
		return
	}
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Bad request: no username provided", http.StatusBadRequest)
		return
	}

	p, err := rt.db.GetAccount(token, token)
	if err != nil {
		rt.internalServerError(err, w)
	}
	if username != p.Username {
		taken, err := rt.db.UsernameTaken(username)
		if err != nil {
			rt.internalServerError(err, w)
			return
		}
		if taken {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		err = rt.db.SetUsername(token, username)
		if err != nil {
			rt.internalServerError(err, w)
			return
		}
	}

	w.Header().Set("content-type", "text/plain")
	_, err = w.Write([]byte(username))
	if err != nil {
		rt.internalServerError(err, w)
	}
}
