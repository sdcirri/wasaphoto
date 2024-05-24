package api

import (
    "github.com/sdgondola/wasaphoto/service/database"
	"github.com/julienschmidt/httprouter"
    "io/ioutil"
	"net/http"
    "strconv"
)

func (rt *_router) commentPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
    textRaw, err := ioutil.ReadAll(r.Body)
    if err != nil {
    	http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
    	return
    }
    text := string(textRaw[:])
    if text == "" {
        http.Error(w, "Bad request: empty comment", http.StatusBadRequest)
        return
    }
    if len(text) > 2048 {
        http.Error(w, "Bad request: comment too long", http.StatusBadRequest)
        return
    }

    cID, err := rt.db.CommentPost(id, postID, text)
    if err == database.ErrUserNotFound {
        // This is kinda suspicious, likely a forged cookie
        http.Error(w, "Bad request: hacking attempt?!", http.StatusBadRequest)
    } else if err == database.ErrPostNotFound {
        http.Error(w, database.ErrPostNotFound.Error(), http.StatusNotFound)
    } else if err == database.ErrUserIsBlocked {
        http.Error(w, "Cannot like post: user blocked you!", http.StatusForbidden)
    } else if err != nil {
        http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
    } else {
        w.Header().Set("content-type", "text/plain")
        w.Write([]byte(string(cID)))
    }
}
