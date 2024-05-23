package database

import (
	"github.com/sdgondola/wasaphoto/service/globaltime"
	"encoding/base64"
	"image/jpeg"
	"image"
	"bytes"
	"os"
)

func (db *appdbimpl) NewPost(op string, imgB64 string, caption string) (int64, error) {
	exists, err := db.UserExists(op)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, ErrUserNotFound
	}
	rawImg, err := base64.StdEncoding.DecodeString(imgB64)
    if err != nil {
    	return 0, err
    }
    img, _, err := image.Decode(bytes.NewReader(rawImg))
    if err != nil {
        return 0, ErrBadImage
    }

	// We first need to insert a stub post in order to get the postID we need for storage
	ins, err := db.c.Prepare("insert into Posts values (?, ?, ?, ?) returning postID")
	if err != nil {
		return 0, err
	}
	del, err := db.c.Prepare("delete from Posts where postID = ?")
	if err != nil {
		return 0, err
	}
	res, err := ins.Exec("", globaltime.Now(), op, caption)
	if err != nil {
		postID, err2 := res.LastInsertId()
		if err2 == nil {
			del.Exec(postID)
		}
		return 0, err
	}
	postID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	imgPath := "/srv/wasaphoto/posts/" + string(postID) + ".jpg"
	imgFile, err := os.Create(imgPath)
	if err != nil {
		return 0, err
	}
	defer imgFile.Close()
	jpegOptions := &jpeg.Options{Quality: 85}
	err = jpeg.Encode(imgFile, img, jpegOptions)
	if err != nil {
		del.Exec(postID)
		return 0, err
	}
	insImg, err := db.c.Prepare("update Posts set img_path = ? where postID = ?")
	if err != nil {
		del.Exec(postID)
		return 0, err
	}
	_, err = insImg.Exec(imgPath, postID)
	if err != nil {
		del.Exec(postID)
		return 0, err
	}
	return postID, nil
}
