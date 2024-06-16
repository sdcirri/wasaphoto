package database

import "strings"

func (db *appdbimpl) UserExists(userID int64) (bool, error) {
	var ok bool
	err := db.c.QueryRow("select exists(select 1 from Users where userID = ?)", userID).Scan(&ok)
	if err != nil {
		return false, err
	}
	return ok, nil
}

func (db *appdbimpl) UsernameTaken(login string) (bool, error) {
	login = strings.ToLower(login)
	var ok bool
	err := db.c.QueryRow("select exists(select 1 from Users where username = ?)", login).Scan(&ok)
	if err != nil {
		return false, err
	}
	return ok, nil
}
