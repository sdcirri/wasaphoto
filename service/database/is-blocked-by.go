package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) IsBlockedBy(blocked int64, blocker int64) (bool, error) {
	exist, err := db.UsersExist(blocked, blocker)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, ErrUserNotFound
	}
	var ok bool
	err = db.c.QueryRow("select exists(select 1 from Blocks where blocker = ? and blocked = ?)", blocker, blocked).Scan(&ok)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false, err
	}
	return ok, nil
}
