package database

import (
	"os"
	"strconv"
)

func (db *appdbimpl) RmPost(op int64, post int64) error {
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

	selc, err := db.c.Query("select commentID from Comments where post = ?", post)
	if err != nil {
		return err
	}
	for selc.Next() {
		var c int64
		err = selc.Err()
		if err != nil {
			return err
		}
		err := selc.Scan(&c)
		if err != nil {
			return err
		}
		err = db.DeleteComment(op, c)
		if err != nil {
			return err
		}
	}

	deltran, err := db.c.Begin()
	if err != nil {
		err2 := deltran.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}

	_, err = deltran.Exec("delete from LikesP where post = ?", post)
	if err != nil {
		err2 := deltran.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}

	_, err = deltran.Exec("delete from Posts where postID = ?", post)
	if err != nil {
		err2 := deltran.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}

	err = deltran.Commit()
	if err != nil {
		return err
	}

	postPath := db.installRoot + "/" + strconv.FormatInt(op, 10) + "/posts/" + strconv.FormatInt(post, 10) + ".jpg"
	return os.Remove(postPath)
}
