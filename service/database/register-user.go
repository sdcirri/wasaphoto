package database

import (
	"os"
	"strings"
)

func getAllowedCharset() string {
	chrset := ".-_"
	for i := 'a'; i <= 'z'; i++ {
		chrset += string(i)
	}
	for i := '0'; i <= 9; i++ {
		chrset += string(i)
	}
	return chrset
}

func (db *appdbimpl) RegisterUser(username string) error {
	allowed_charset := getAllowedCharset()
	username = strings.ToLower(username)

	for _, c := range username {
		if !strings.Contains(allowed_charset, string(c)) {
			return ErrBadCharset
		}
	}

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
