package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) DeleteComment(user int64, commentID int64) error {
	exists, err := db.UserExists(user)
	if err != nil {
		return err
	}
	if !exists {
		return ErrUserNotFound
	}

	var oc, op int64
	err = db.c.QueryRow("select author from Comments where commentID = ?", commentID).Scan(&oc)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrCommentNotFound
	} else if err != nil {
		return err
	}
	err = db.c.QueryRow("select p.author from Posts p, Comments c where c.post = p.postID and c.commentID = ?", commentID).Scan(&op)
	if err != nil {
		return err
	}
	if user != oc && user != op { // OP should be able to delete comments on its own post
		return ErrUserIsNotAuthor
	}
	deltran, err := db.c.Begin()
	if err != nil {
		return err
	}
	_, err = deltran.Exec("delete from Comments where commentID = ?", commentID)
	if err != nil {
		err2 := deltran.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}
	_, err = deltran.Exec("delete from LikesC where comment = ?", commentID)
	if err != nil {
		err2 := deltran.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}

	return deltran.Commit()
}
