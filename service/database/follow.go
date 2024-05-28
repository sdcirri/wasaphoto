package database

import "strings"

func (db *appdbimpl) Follow(follower string, toFollow string) error {
	follower = strings.ToLower(follower)
	toFollow = strings.ToLower(toFollow)
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
