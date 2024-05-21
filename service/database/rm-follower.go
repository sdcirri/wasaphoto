package database

func (db *appdbimpl) RmFollower(user string, follower string) error {
	return db.Unfollow(follower, user)
}

