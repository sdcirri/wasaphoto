package api

import (
    "github.com/sdgondola/wasaphoto/service/database"
	"github.com/julienschmidt/httprouter"
    "encoding/json"
	"net/http"
    "strconv"
)

// Here we don't want to return the likes explicitly, but rather the like count
type PostFE struct {
	PostID		int64		`json:"postID"`
	ImageB64	string		`json:"imageB64"`
	PubTime		string		`json:"pubTime"`
	Caption		string		`json:"caption"`
	Author		string		`json:"author"`
	LikeCount	int     	`json:"likeCount"`
	Comments	[]int64		`json:"comments"`
}

func post2postFE(post database.Post) PostFE {
    return PostFE {
        PostID: post.PostID,
        ImageB64: post.ImageB64,
        PubTime: post.PubTime,
        Caption: post.Caption,
        Author: post.Author,
        LikeCount: len(post.Likes),
        Comments: post.Comments,
    }
}

func (rt *_router) getPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
        j, err := json.Marshal(post2postFE(post))
        if err != nil {
            http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("content-type", "application/json")
        w.Write(j)
    }
}
