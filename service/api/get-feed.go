package api

import (
    "github.com/sdgondola/wasaphoto/service/database"
	"github.com/julienschmidt/httprouter"
    "encoding/json"
	"net/http"
)

func (rt *_router) getFeed(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    idc, err := r.Cookie("WASASESSIONID")
    if err == http.ErrNoCookie {
        http.Error(w, "Unauthenticated", http.StatusUnauthorized)
        return
    } else if err != nil {
    	http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
    	return
    }
    id := idc.Value

    feed, err := rt.db.GetFeed(id)
    if err == database.ErrUserNotFound {
        http.Error(w, "Bad request: hacking attempt?!", http.StatusBadRequest)
    } else if err != nil {
        http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
    } else {
        j, err := json.Marshal(feed)
        if err != nil {
        	http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
        	return
        }
        w.Header().Set("content-type", "application/json")
        w.Write(j)
    }
}
