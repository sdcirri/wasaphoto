package database

func (db *appdbimpl) Follow(follower int64, toFollow int64) error {
	exist, err := db.UsersExist(follower, toFollow)
	if err != nil {
		return err
	}
	if !exist {
		return ErrUserNotFound
	}
	blocked, err := db.IsBlockedBy(follower, toFollow)
	if err != nil {
		return err
	}
	if blocked {
		return ErrUserIsBlocked
	}
	already, err := db.Follows(follower, toFollow)
	if err != nil {
		return err
	}
	if already {
		return ErrAlreadyFollowing
	}
	ins, err := db.c.Prepare("insert into Follows values (?, ?)")
	if err != nil {
		return err
	}
	_, err = ins.Exec(follower, toFollow)
	return err
}
