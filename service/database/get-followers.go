package database

import "strings"

func (db *appdbimpl) GetFollowers(id string) ([]string, error) {
	id = strings.ToLower(id)
	followers := make([]string, 0)
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
		var f string
		err = q.Scan(&f)
		if err != nil {
			return followers, err
		}
		followers = append(followers, f)
	}
	return followers, nil
}
