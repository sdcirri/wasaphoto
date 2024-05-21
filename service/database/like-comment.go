package database

func (db *appdbimpl) LikeComment(user string, commentID int64) error {
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
	ins, err := db.c.Prepare("insert into LikesC values (?, ?)")
	if err != nil {
		return err
	}
	_, err = ins.Exec(user, commentID)
	return err
}
