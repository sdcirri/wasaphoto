package database

import (
	"os"
	"strings"
)

func (db *appdbimpl) RegisterUser(username string) error {
	username = strings.ToLower(username)
	exists, err := db.UserExists(username)
	if err != nil {
		return err
	}
	if exists {
		return ErrUserAlreadyExists
	}
	ins, err := db.c.Prepare("insert into Users(username) values (?)")
	if err != nil {
		return err
	}
	_, err = ins.Exec(username)
	if err != nil {
		return err
	}
	// Directory to store user data such as profile picture and posted pictures
	err = os.MkdirAll("/srv/wasaphoto/"+username, 0755)
	return err
}
