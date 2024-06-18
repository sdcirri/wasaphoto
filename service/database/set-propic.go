package database

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"log"
	"os"
	"strconv"
)

func (db *appdbimpl) SetProPic(userID int64, imgB64 string) error {
	exists, err := db.UserExists(userID)
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
	dstPath := "/srv/wasaphoto/" + strconv.FormatInt(userID, 10) + "/propic.jpg"
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()
	log.Println("Created file", dstPath)

	jpegOptions := &jpeg.Options{Quality: 85}
	err = jpeg.Encode(dst, img, jpegOptions)
	log.Println("Encoded image")
	if err != nil {
		return err
	}
	tran, err := db.c.Prepare("update Users set propic = ? where userID = ?")
	if err != nil {
		return err
	}
	_, err = tran.Exec(dstPath, userID)
	log.Println("Profile picture updated")
	return err
}
