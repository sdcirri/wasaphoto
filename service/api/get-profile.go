package api

import (
    "github.com/sdgondola/wasaphoto/service/database"
	"github.com/julienschmidt/httprouter"
    "encoding/base64"
    "encoding/json"
    "io/ioutil"
	"net/http"
)

// This is slightly different from the profile returned by the database,
// as we do not want to return the image path (obviously), but rather the image
// in base64 format
type Profile struct {
	Username	string		`json:"username"`
	ProPic     	string		`json:"proPic"`
	Followers	uint		`json:"followers"`
	Following	uint		`json:"following"`
	Posts		[]int64		`json:"posts"`
}

func DBAcc2FEAcc(acc database.Account) (Profile, error) {
    p := Profile{
        Username: acc.Username,
        Followers: acc.Followers,
        Following: acc.Following,
        Posts: acc.Posts,
    }
    img, err := ioutil.ReadFile(acc.ProPicPath)
    if err != nil {
        return p, err
    }
    p.ProPic = base64.StdEncoding.EncodeToString(img)
    return p, nil
}

func (rt *_router) getProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    var canViewPosts bool
    var id string
    idc, err := r.Cookie("WASASESSIONID")
    if err == http.ErrNoCookie {
        canViewPosts = false
        id = ""
    } else if err != nil {
    	http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
    	return
    } else {
        id = idc.Value
        exists, err := rt.db.UserExists(id)
        if err != nil {
            http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
            return
        }
        if !exists {
            // This is kinda suspicious, likely a forged cookie
            http.Error(w, "Bad request: hacking attempt?!", http.StatusBadRequest)
            return
        }
    }

    toView := ps.ByName("username")
    profile, err := rt.db.GetAccount(id, toView)
    if err == database.ErrUserIsBlocked {
        http.Error(w, "Forbidden: user blocked you!", http.StatusForbidden)
    } else if err == database.ErrUserNotFound {
        http.Error(w, "User not found", http.StatusNotFound)
    } else if err != nil {
    	http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
    } else {
        if !canViewPosts {
            profile.Posts = make([]int64, 0)
        }
        ret, err := DBAcc2FEAcc(profile)
        if err != nil {
        	http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
        	return
        }
        j, err := json.Marshal(ret)
        if err != nil {
        	http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
        	return
        }
        w.Header().Set("content-type", "application/json")
        w.Write(j)
    }
}
