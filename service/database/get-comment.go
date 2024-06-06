package database

import (
	"database/sql"
	"errors"
	"time"
)

func (db *appdbimpl) GetComment(commentID int64) (Comment, error) {
	var c Comment
	var pubts time.Time
	err := db.c.QueryRow("select * from Comments where commentID = ?", commentID).Scan(&c.CommentID, &pubts, &c.Author, &c.PostID, &c.Content)
	if errors.Is(err, sql.ErrNoRows) {
		return c, ErrCommentNotFound
	} else if err != nil {
		return c, err
	}
	c.Time = pubts.Format(time.RFC3339)
	err = db.c.QueryRow("select count(*) from LikesC where comment = ?", commentID).Scan(&c.Likes)
	if errors.Is(err, sql.ErrNoRows) {
		c.Likes = 0
	} else if err != nil {
		return c, err
	}

	return c, nil
}
