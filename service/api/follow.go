package api

import (
    "github.com/sdgondola/wasaphoto/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) follow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    id, err := r.Cookie("WASASESSIONID")
    if err == http.ErrNoCookie {
        http.Error(w, "Unauthenticated", http.StatusUnauthorized)
        return
    } else if err != nil {
    	http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
    	return
    }
    toFollow := ps.ByName("username")
    if toFollow == "" {
        http.Error(w, "Bad request: no username provided", http.StatusBadRequest)
        return
    }
    if toFollow == id.Value {
        http.Error(w, "Bad request: you cannot follow yourself!", http.StatusBadRequest)
        return
    }
    err = rt.db.Follow(id.Value, toFollow)
    if err == database.ErrUserNotFound {
        http.Error(w, "Bad request: no such user", http.StatusBadRequest)
    } else if err == database.ErrUserIsBlocked {
        http.Error(w, "Forbidden: user blocked you!", http.StatusForbidden)
    } else if err == database.ErrAlreadyFollowing {
        http.Error(w, "Bad request: already following", http.StatusBadRequest)
    } else if err != nil {
    	http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
    } else {
        w.WriteHeader(http.StatusCreated)
    }
}
