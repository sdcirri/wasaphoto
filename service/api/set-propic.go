package api

import (
	"github.com/sdgondola/wasaphoto/service/database"
	"github.com/julienschmidt/httprouter"
    "image/jpeg"
    "image"
	"net/http"
    "os"
)

func (rt *_router) setProPic(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    id, err := r.Cookie("WASASESSIONID")
    if err == http.ErrNoCookie {
        http.Error(w, "Unauthenticated", http.StatusUnauthorized)
        return
    } else if err != nil {
    	http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
    	return
    }

    // Get image
    err = r.ParseMultipartForm(10 << 20)
    if err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }
    ffile, _, err := r.FormFile("image")
    if err != nil {
        http.Error(w, "Bad image", http.StatusBadRequest)
        return
    }
    defer ffile.Close()
    // Convert to jpeg for storage
    dstPath := "/srv/wasaphoto/" + id.Value + "/propic.jpg"
    dst, err := os.Create(dstPath)
    if err != nil {
        http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
        return
    }
    defer dst.Close()
    propic, _, err := image.Decode(ffile)
    if err != nil {
        http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
        return
    }
    jpegOptions := &jpeg.Options{Quality: 85}
    err = jpeg.Encode(dst, propic, jpegOptions)
    if err != nil {
        http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
        return
    }
    // Write to database
    err = rt.db.SetProPic(id.Value, dstPath)
	if err == database.ErrUserNotFound {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
    if err != nil {
        http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}
