package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UserInfo struct {
	Username string `json:"name"`
}

func (rt *_router) login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// First check if the user is already logged in
	userID, err := rt.getAuthToken(r)
	if errors.Is(err, ErrNoAuth) {
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
		Value: strconv.FormatInt(userID, 10),
		Path:  "/",
	})
	w.Header().Set("content-type", "text-plain")
	_, err = w.Write([]byte(strconv.FormatInt(userID, 10)))
	if err != nil {
		rt.internalServerError(err, w)
		return
	}
}
