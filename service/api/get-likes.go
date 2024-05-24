package api

import (
    "github.com/sdgondola/wasaphoto/service/database"
	"github.com/julienschmidt/httprouter"
    "encoding/json"
	"net/http"
    "strconv"
)

func (rt *_router) getLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
    post, err := rt.db.GetPost(id, postID)
    if err == database.ErrUserNotFound {
        // This is kinda suspicious, likely a forged cookie
        http.Error(w, "Bad request: hacking attempt?!", http.StatusBadRequest)
    } else if err == database.ErrPostNotFound {
        http.Error(w, database.ErrPostNotFound.Error(), http.StatusNotFound)
    } else if err == database.ErrUserIsBlocked {
        http.Error(w, "Cannot view post: user blocked you!", http.StatusForbidden)
    } else if err != nil {
        http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
    } else {
        if id != post.Author {
            http.Error(w, "Forbidden: you can't view somebody else's likes", http.StatusForbidden)
            return
        }
        j, err := json.Marshal(post.Likes)
        if err != nil {
            http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("content-type", "application/json")
        w.Write(j)
    }
}
