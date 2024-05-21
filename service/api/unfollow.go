package api

import (
    "github.com/sdgondola/wasaphoto/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) unfollow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    id, err := r.Cookie("WASASESSIONID")
    if err == http.ErrNoCookie {
        http.Error(w, "Unauthenticated", http.StatusUnauthorized)
        return
    } else if err != nil {
    	http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
    	return
    }
    toUnfollow := ps.ByName("username")
    if toUnfollow == "" {
        http.Error(w, "Bad request: no username provided", http.StatusBadRequest)
        return
    }
    if toUnfollow == id.Value {
        http.Error(w, "Bad request: you cannot unfollow yourself!", http.StatusBadRequest)
        return
    }
    err = rt.db.Unfollow(id.Value, toUnfollow)
    if err == database.ErrUserNotFound {
        http.Error(w, "Bad request: no such user", http.StatusBadRequest)
    } else if err == database.ErrUserIsBlocked {
        http.Error(w, "Forbidden: user blocked you!", http.StatusForbidden)
    } else if err == database.ErrNotFollowing {
        http.Error(w, "Bad request: not following", http.StatusBadRequest)
    } else if err != nil {
    	http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
    } else {
        w.WriteHeader(http.StatusCreated)
    }
}
