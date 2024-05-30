package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/sdgondola/wasaphoto/service/database"
)

var ErrNoAuth = errors.New("unauthenticated")

func (rt *_router) getAuthToken(r *http.Request) (string, error) {
	tokenSplit := strings.Split(strings.ToLower(r.Header.Get("Authorization")), "bearer")
	if len(tokenSplit) < 2 {
		return "", ErrNoAuth
	}
	token := strings.TrimSpace(tokenSplit[1])
	valid, err := rt.db.UserExists(token)
	if err != nil {
		return token, err
	}
	if !valid {
		return token, database.ErrUserNotFound
	}
	return token, nil
}
