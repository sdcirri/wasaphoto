package api

import (
    "github.com/sdgondola/wasaphoto/service/database"
	"github.com/julienschmidt/httprouter"
    "encoding/json"
	"net/http"
    "strconv"
)

func (rt *_router) getComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    idc, err := r.Cookie("WASASESSIONID")
    if err == http.ErrNoCookie {
        http.Error(w, "Unauthenticated", http.StatusUnauthorized)
        return
    } else if err != nil {
    	http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
    	return
    }
    id := idc.Value
    commentID, err := strconv.ParseInt(ps.ByName("commentID"), 10, 64)
    if err != nil {
        http.Error(w, database.ErrPostNotFound.Error(), http.StatusNotFound)
        return
    }
    comment, err := rt.db.GetComment(commentID)
    if err == database.ErrUserNotFound {
        // This is kinda suspicious, likely a forged cookie
        http.Error(w, "Bad request: hacking attempt?!", http.StatusBadRequest)
    } else if err == database.ErrCommentNotFound {
        http.Error(w, database.ErrCommentNotFound.Error(), http.StatusNotFound)
    } else if err != nil {
        http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
    } else {
        blocked, err := rt.db.IsBlockedBy(id, comment.Author)
        if err != nil {
            http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
            return
        }
        if blocked {
            http.Error(w, "Cannot view comment: user blocked you!", http.StatusForbidden)
            return
        }
        j, err := json.Marshal(comment)
        if err != nil {
            http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("content-type", "application/json")
        w.Write(j)
    }
}
