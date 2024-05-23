package database

import (
	"encoding/base64"
	"image/jpeg"
	"image"
	"bytes"
	"os"
)

func (db *appdbimpl) SetProPic(username string, imgB64 string) error {
	exists, err := db.UserExists(username)
	if err != nil {
		return err
	}
	if !exists {
		return ErrUserNotFound
	}
	rawImg, err := base64.StdEncoding.DecodeString(imgB64)
    if err != nil {
    	return err
    }
    img, _, err := image.Decode(bytes.NewReader(rawImg))
    if err != nil {
        return ErrBadImage
    }
	dstPath := "/srv/wasaphoto/" + username + "/propic.jpg"
	dst, err := os.Create(dstPath)
	if err != nil {
        return err
    }
	defer dst.Close()

    jpegOptions := &jpeg.Options{Quality: 85}
    err = jpeg.Encode(dst, img, jpegOptions)
    if err != nil {
        return err
    }
	tran, err := db.c.Prepare("update table Users set propic = ? where username = ?")
	if err != nil {
		return err
	}
	_, err = tran.Exec(dstPath, username)
	return err
}
