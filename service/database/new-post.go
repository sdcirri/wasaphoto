package database

import (
	"github.com/sdgondola/wasaphoto/service/globaltime"
	"os"
)

func (db *appdbimpl) NewPost(op string, imgpath string, caption string) (int64, error) {
	_, err := os.Stat(imgpath)
	if err != nil {
		return 0, err
	}
	exists, err := db.UserExists(op)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, ErrUserNotFound
	}
	ins, err := db.c.Prepare("insert into Posts values (?, ?, ?, ?) returning postID")
	if err != nil {
		return 0, err
	}
	res, err := ins.Exec(imgpath, globaltime.Now(), op, caption)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
