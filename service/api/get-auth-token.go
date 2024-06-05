package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/sdgondola/wasaphoto/service/database"
)

var ErrNoAuth = errors.New("unauthenticated")

func (rt *_router) getAuthToken(r *http.Request) (int64, error) {
	tokenSplit := strings.Split(strings.ToLower(r.Header.Get("Authorization")), "bearer")
	if len(tokenSplit) < 2 {
		return 0, ErrNoAuth
	}
	token, err := strconv.ParseInt(strings.TrimSpace(tokenSplit[1]), 10, 64)
	if err != nil {
		return token, err
	}
	valid, err := rt.db.UserExists(token)
	if err != nil {
		return token, err
	}
	if !valid {
		return token, database.ErrUserNotFound
	}
	return token, nil
}
