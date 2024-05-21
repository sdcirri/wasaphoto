package database

import "database/sql"

func (db *appdbimpl) UserExists(login string) (bool, error) {
	var ok bool
	err := db.c.QueryRow("select exists(select 1 from Users where username = ?)", login).Scan(&ok)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return ok, nil
}
