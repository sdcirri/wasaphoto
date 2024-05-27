package database

import (
	"os"
	"strconv"
)

func (db *appdbimpl) RmPost(op string, post int64) error {
	exists, err := db.PostExists(post)
	if err != nil {
		return err
	}
	if !exists {
		return ErrPostNotFound
	}
	record, err := db.GetPost(op, post)
	if err != nil {
		return err
	}
	if op != record.Author {
		return ErrUserIsNotAuthor
	}

	del, err := db.c.Prepare("delete from Posts where postID = ?")
	if err != nil {
		return err
	}
	_, err = del.Exec(post)
	if err != nil {
		return err
	}
	err = os.Remove("/srv/wasaphoto/posts/" + strconv.FormatInt(post, 10) + ".jpg")
	return err
}
