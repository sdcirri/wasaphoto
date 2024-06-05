package database

func (db *appdbimpl) GetFollowers(id int64) ([]int64, error) {
	followers := make([]int64, 0)
	exists, err := db.UserExists(id)
	if err != nil {
		return followers, err
	}
	if !exists {
		return followers, ErrUserNotFound
	}
	q, err := db.c.Query("select follower from Follows where following = ?", id)
	if err != nil {
		return followers, err
	}
	for q.Next() {
		var f int64
		err = q.Scan(&f)
		if err != nil {
			return followers, err
		}
		followers = append(followers, f)
	}
	return followers, nil
}
