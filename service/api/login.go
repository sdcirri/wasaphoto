package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sdgondola/wasaphoto/service/database"
)

type UserInfo struct {
	Username string `json:"name"`
}

func (rt *_router) login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// First check if the user is already logged in
	userID, err := rt.getAuthToken(r)
	if errors.Is(err, ErrNoAuth) {
		http.Error(w, "Bad request: no auth token provided", http.StatusBadRequest)
		return
	} else if errors.Is(err, database.ErrUserNotFound) {
		if len(userID) < 3 || len(userID) > 127 {
			http.Error(w, "Bad userID", http.StatusBadRequest)
			return
		}
		var info UserInfo
		err = json.NewDecoder(r.Body).Decode(&info)
		if err != nil {
			rt.internalServerError(err, w)
		}

		userID, err = rt.db.RegisterUser(info.Username)
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
		Value: userID,
		Path:  "/",
	})
	w.Header().Set("content-type", "text-plain")
	_, err = w.Write([]byte(userID))
	if err != nil {
		rt.internalServerError(err, w)
		return
	}
}
