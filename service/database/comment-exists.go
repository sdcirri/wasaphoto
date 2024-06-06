package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) CommentExists(commentID int64) (bool, error) {
	var ok bool
	err := db.c.QueryRow("select exists(select 1 from Comments where commentID = ?)", commentID).Scan(&ok)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false, err
	}
	return ok, nil
}
