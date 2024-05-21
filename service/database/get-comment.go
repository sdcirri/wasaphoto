package database

import "database/sql"

func (db *appdbimpl) GetComment(commentID int64) (Comment, error) {
	var c Comment
	err := db.c.QueryRow("select * from Comments where commentID = ?", commentID).Scan(&c.CommentID, &c.Time, &c.Author, &c.PostID, &c.Content)
	if err == sql.ErrNoRows {
		return c, ErrCommentNotFound
	} else if err != nil {
		return c, err
	}
	err = db.c.QueryRow("select count(*) from LikesC where comment = ?", commentID).Scan(&c.Likes)
	if err == sql.ErrNoRows {
		c.Likes = 0
	} else if err != nil {
		return c, err
	}

	return c, nil
}
