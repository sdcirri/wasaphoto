package database

import (
	"os"
	"strconv"
	"strings"
)

func (db *appdbimpl) RegisterUser(username string) (int64, error) {
	username = strings.ToLower(username)
	exists, err := db.UsernameTaken(username)
	if err != nil {
		return 0, err
	}
	if exists {
		return 0, ErrUserAlreadyExists
	}

	ins, err := db.c.Prepare("insert into Users(username, propic) values (?, ?)")
	if err != nil {
		return 0, err
	}
	res, err := ins.Exec(username, db.installRoot+"/propic_default.jpg")
	if err != nil {
		return 0, err
	}
	userID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	// Directory to store user data such as profile picture and posted pictures
	err = os.MkdirAll(db.installRoot+"/"+strconv.FormatInt(userID, 10), 0755)
	return userID, err
}
