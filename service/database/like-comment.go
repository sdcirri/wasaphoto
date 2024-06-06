package database

import (
	"database/sql"
	"errors"

	"github.com/mattn/go-sqlite3"
)

func (db *appdbimpl) LikeComment(user int64, commentID int64) error {
	exists, err := db.UserExists(user)
	if err != nil {
		return err
	}
	if !exists {
		return ErrUserNotFound
	}

	var oc int64
	err = db.c.QueryRow("select author from Comments where commentID = ?", commentID).Scan(&oc)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrCommentNotFound
	} else if err != nil {
		return err
	}
	blocked, err := db.IsBlockedBy(user, oc)
	if err != nil {
		return err
	}
	if blocked {
		return ErrUserIsBlocked
	}
	ins, err := db.c.Prepare("insert into LikesC values (?, ?)")
	if err != nil {
		return err
	}
	_, err = ins.Exec(user, commentID)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.Code == sqlite3.ErrConstraint {
			return ErrAlreadyLiked
		}
		return err
	}
	return nil
}
