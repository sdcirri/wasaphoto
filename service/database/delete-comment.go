package database

import (
	"database/sql"
	"errors"
	"strings"
)

func (db *appdbimpl) DeleteComment(user string, commentID int64) error {
	user = strings.ToLower(user)
	exists, err := db.UserExists(user)
	if err != nil {
		return err
	}
	if !exists {
		return ErrUserNotFound
	}

	var oc string
	err = db.c.QueryRow("select author from Comments where commentID = ?", commentID).Scan(&oc)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrCommentNotFound
	} else if err != nil {
		return err
	}
	if oc != user {
		return ErrUserIsNotAuthor
	}

	del, err := db.c.Prepare("delete from Comments where commentID = ?")
	if err != nil {
		return err
	}
	_, err = del.Exec(user, commentID)
	return err
}
