package database

func (db *appdbimpl) LikePost(user string, postID int64) error {
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
	return err
}
