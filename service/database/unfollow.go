package database

func (db *appdbimpl) Unfollow(follower int64, toUnfollow int64) error {
	exist, err := db.UsersExist(follower, toUnfollow)
	if err != nil {
		return err
	}
	if !exist {
		return ErrUserNotFound
	}
	follows, err := db.Follows(follower, toUnfollow)
	if err != nil {
		return err
	}
	if !follows {
		return ErrNotFollowing
	}
	del, err := db.c.Prepare("delete from Follows where follower = ?, and following = ?")
	if err != nil {
		return err
	}
	_, err = del.Exec(follower, toUnfollow)
	return err
}
