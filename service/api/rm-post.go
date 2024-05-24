package api

import (
    "github.com/sdgondola/wasaphoto/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
    "strconv"
)

func (rt *_router) rmPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    idc, err := r.Cookie("WASASESSIONID")
    if err == http.ErrNoCookie {
        http.Error(w, "Unauthenticated", http.StatusUnauthorized)
        return
    } else if err != nil {
    	http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
    	return
    }
    id := idc.Value
    postID, err := strconv.ParseInt(ps.ByName("postID"), 10, 64)
    if err != nil {
        http.Error(w, database.ErrPostNotFound.Error(), http.StatusNotFound)
        return
    }

    err = rt.db.RmPost(id, postID)
    if err == database.ErrPostNotFound {
        http.Error(w, database.ErrPostNotFound.Error(), http.StatusNotFound)
    } else if err == database.ErrUserIsNotAuthor {
        http.Error(w, database.ErrUserIsNotAuthor.Error(), http.StatusForbidden)
    } else if err != nil {
        http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
    } else {
        w.WriteHeader(http.StatusOK)
    }
}
