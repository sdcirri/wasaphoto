package database

import (
	"errors"

	"github.com/mattn/go-sqlite3"
)

func (db *appdbimpl) SetUsername(userID int64, username string) error {
	up, err := db.c.Prepare("update Users set username = ? where userID = ?")
	if err != nil {
		return err
	}
	_, err = up.Exec(username, userID)
	var sqliteErr sqlite3.Error
	if errors.As(err, &sqliteErr) && sqliteErr.Code == sqlite3.ErrConstraint {
		return ErrUsernameAlreadyTaken
	}
	return nil
}
