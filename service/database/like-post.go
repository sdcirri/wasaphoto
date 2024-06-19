package database

import (
	"errors"

	"github.com/mattn/go-sqlite3"
)

func (db *appdbimpl) LikePost(user int64, postID int64) error {
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
	p, err := db.GetPost(user, postID)
	if err != nil {
		return err
	}
	blocked, err := db.IsBlockedBy(user, p.Author)
	if err != nil {
		return err
	}
	if blocked {
		return ErrUserIsBlocked
	}

	ins, err := db.c.Prepare("insert into LikesP values (?, ?)")
	if err != nil {
		return err
	}
	_, err = ins.Exec(user, postID)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && sqliteErr.Code == sqlite3.ErrConstraint {
			return ErrAlreadyLiked
		}
		return err
	}
	return nil
}
