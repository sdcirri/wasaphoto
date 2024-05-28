package database

import (
	"database/sql"
	"strings"
)

func (db *appdbimpl) Follows(follower string, following string) (bool, error) {
	follower = strings.ToLower(follower)
	following = strings.ToLower(following)
	exist, err := db.UsersExist(follower, following)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, ErrUserNotFound
	}
	var ok bool
	err = db.c.QueryRow("select exists(select 1 from Follows where follower = ? and following = ?)", follower, following).Scan(&ok)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return ok, nil
}
