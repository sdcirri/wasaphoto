package database

func (db *appdbimpl) RmFollower(user int64, follower int64) error {
	return db.Unfollow(follower, user)
}
