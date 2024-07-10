package database

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
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

	image.RegisterFormat("jpg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("gif", "gif", gif.Decode, gif.DecodeConfig)

	rawImg, err := base64.StdEncoding.DecodeString(imgB64)
	if err != nil {
		return err
	}
	img, _, err := image.Decode(bytes.NewReader(rawImg))
	if err != nil {
		return ErrBadImage
	}
	dstPath := db.installRoot + "/" + strconv.FormatInt(userID, 10) + "/propic.jpg"
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
	tran, err := db.c.Prepare("update Users set propic = ? where userID = ?")
	if err != nil {
		return err
	}
	_, err = tran.Exec(dstPath, userID)
	return err
}
