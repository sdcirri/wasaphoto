package api

import (
    "github.com/sdgondola/wasaphoto/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
    "strconv"
)

func (rt *_router) unlikePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    idc, err := r.Cookie("WASASESSIONID")
    if err == http.ErrNoCookie {
        http.Error(w, "Unauthenticated", http.StatusUnauthorized)
        return
    } else if err != nil {
        http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
        return
    }
    toBlock := ps.ByName("postID")
    if toBlock == "" {
        http.Error(w, "Bad request: no username provided", http.StatusBadRequest)
        return
    }
    id := idc.Value
    postID, err := strconv.ParseInt(ps.ByName("postID"), 10, 64)
    if err != nil {
        http.Error(w, database.ErrPostNotFound.Error(), http.StatusNotFound)
        return
    }

    err = rt.db.UnlikePost(id, postID)
    if err == database.ErrUserNotFound {
        // This is kinda suspicious, likely a forged cookie
        http.Error(w, "Bad request: hacking attempt?!", http.StatusBadRequest)
    } else if err == database.ErrPostNotFound {
        http.Error(w, database.ErrPostNotFound.Error(), http.StatusNotFound)
    } else if err == database.ErrDidNotLike {
        http.Error(w, database.ErrDidNotLike.Error(), http.StatusBadRequest)
    } else if err != nil {
        http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
    } else {
        w.WriteHeader(http.StatusCreated)
    }
}
