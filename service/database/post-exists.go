package database

import "database/sql"

func (db *appdbimpl) PostExists(postID int64) (bool, error) {
	var ok bool
	err := db.c.QueryRow("select exists (select 1 from Posts where postID = ?)", postID).Scan(&ok)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return ok, nil
}
