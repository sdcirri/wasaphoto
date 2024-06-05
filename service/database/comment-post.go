package database

import (
	"github.com/sdgondola/wasaphoto/service/globaltime"
)

func (db *appdbimpl) CommentPost(user int64, postID int64, comment string) (int64, error) {
	exists, err := db.UserExists(user)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, ErrUserNotFound
	}
	exists, err = db.PostExists(postID)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, ErrPostNotFound
	}
	var op int64
	err = db.c.QueryRow("select author from Posts where postID = ?", postID).Scan(&op)
	if err != nil {
		return 0, err
	}
	blocked, err := db.IsBlockedBy(user, op)
	if err != nil {
		return 0, err
	}
	if blocked {
		return 0, ErrUserIsBlocked
	}
	ins, err := db.c.Prepare("insert into Comments values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	res, err := ins.Exec(globaltime.Now(), user, postID, comment)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
