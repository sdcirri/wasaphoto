package api

import (
    "github.com/sdgondola/wasaphoto/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) unblock(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    id, err := r.Cookie("WASASESSIONID")
    if err == http.ErrNoCookie {
        http.Error(w, "Unauthenticated", http.StatusUnauthorized)
        return
    } else if err != nil {
    	http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
    	return
    }
    toUnblock := ps.ByName("username")
    if toUnblock == "" {
        http.Error(w, "Bad request: no username provided", http.StatusBadRequest)
        return
    }
    if toUnblock == id.Value {
        http.Error(w, "Bad request: you cannot block yourself!", http.StatusBadRequest)
        return
    }
    err = rt.db.Unblock(id.Value, toUnblock)
    if err == database.ErrUserNotFound {
        http.Error(w, "Bad request: no such user", http.StatusBadRequest)
    } else if err == database.ErrUserNotBlocked {
        http.Error(w, "Bad request: user not blocked", http.StatusBadRequest)
    } else if err != nil {
    	http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
    } else {
        w.WriteHeader(http.StatusCreated)
    }
}
