package database

import "database/sql"

func (db *appdbimpl) UnlikePost(user string, postID int64) error {
	exists, err := db.UserExists(user)
	if err != nil {
		return err
	}
	if !exists {
		return ErrUserNotFound
	}
	exists, err = db.PostExists(postID)
	if err != nil {
		return err
	}
	if !exists {
		return ErrPostNotFound
	}
	var ok bool
	q := db.c.QueryRow("select exists(select 1 from LikesP where user = ? and post = ?)", user, postID).Scan(&ok)
	if q == sql.ErrNoRows || !ok {
		return ErrDidNotLike
	}
	del, err := db.c.Prepare("delete from LikesP where user = ? and post = ?")
	if err != nil {
		return err
	}
	_, err = del.Exec(user, postID)
	return err
}
