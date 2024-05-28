package database

import "strings"

func (db *appdbimpl) Unblock(user string, toUnblock string) error {
	user = strings.ToLower(user)
	toUnblock = strings.ToLower(toUnblock)
	exist, err := db.UsersExist(user, toUnblock)
	if err != nil {
		return err
	}
	if !exist {
		return ErrUserNotFound
	}
	blocked, err := db.IsBlockedBy(toUnblock, user)
	if err != nil {
		return err
	}
	if !blocked {
		return ErrUserNotBlocked
	}
	del, err := db.c.Prepare("delete from Blocks where blocker = ? and blocked = ?")
	if err != nil {
		return err
	}
	_, err = del.Exec(user, toUnblock)
	return err
}
