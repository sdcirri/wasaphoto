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
	ins, err := db.c.Prepare("insert into Users values (?, ?)")
	if err != nil {
		return err
	}
	_, err = ins.Exec(username, db.installRoot+"/propic_default.jpg")
	if err != nil {
		return err
	}
	// Directory to store user data such as profile picture and posted pictures
	err = os.MkdirAll(db.installRoot+"/"+username, 0755)
	return err
}
