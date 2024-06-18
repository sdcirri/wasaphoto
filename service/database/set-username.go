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
	if sqliteErr, ok := err.(sqlite3.Error); ok && errors.Is(sqliteErr.Code, sqlite3.ErrConstraint) {
		return ErrUsernameAlreadyTaken
	}
	return nil
}
