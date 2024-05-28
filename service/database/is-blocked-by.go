package database

import (
	"database/sql"
	"strings"
)

func (db *appdbimpl) IsBlockedBy(blocked string, blocker string) (bool, error) {
	blocked = strings.ToLower(blocked)
	blocker = strings.ToLower(blocker)
	exist, err := db.UsersExist(blocked, blocker)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, ErrUserNotFound
	}
	var ok bool
	err = db.c.QueryRow("select exists(select 1 from Blocks where blocker = ? and blocked = ?)", blocker, blocked).Scan(&ok)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return ok, nil
}
