package database

import "database/sql"

func (db *appdbimpl) RmPost(op string, post int64) error {
	var ok bool
	err := db.c.QueryRow("select * from Posts where postID = ? and author = ?", post, op).Scan(&ok)
	if err == sql.ErrNoRows {
		return ErrPostNotFound
	}
	del, err := db.c.Prepare("delete from Posts where postID = ?")
	if err != nil {
		return err
	}
	_, err = del.Exec(post)
	return err
}
