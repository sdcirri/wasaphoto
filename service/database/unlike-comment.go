package database

import "database/sql"

func (db *appdbimpl) UnlikeComment(user int64, commentID int64) error {
	exists, err := db.UserExists(user)
	if err != nil {
		return err
	}
	if !exists {
		return ErrUserNotFound
	}
	exists, err = db.CommentExists(commentID)
	if err != nil {
		return err
	}
	if !exists {
		return ErrCommentNotFound
	}
	var ok bool
	q := db.c.QueryRow("select exists(select 1 from LikesC where user = ? and comment = ?)", user, commentID).Scan(&ok)
	if q == sql.ErrNoRows || !ok {
		return ErrDidNotLike
	}
	del, err := db.c.Prepare("delete from LikesC where user = ? and comment = ?")
	if err != nil {
		return err
	}
	_, err = del.Exec(user, commentID)
	return err
}
