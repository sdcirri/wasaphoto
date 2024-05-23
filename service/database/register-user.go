package database

import "os"

func (db *appdbimpl) RegisterUser(username string) error {
	exists, err := db.UserExists(username)
	if err != nil {
		return err
	}
	if exists {
		return ErrUserAlreadyExists
	}
	ins, err := db.c.Prepare("insert into Users values (?, '/srv/wasaphoto/propic_default.jpg')")
	if err != nil {
		return err
	}
	_, err = ins.Exec(username)
	if err != nil {
		return err
	}
	// Directory to store user data such as profile picture and posted pictures
	err = os.MkdirAll("/srv/wasaphoto/" + username, 0755)
	return err
}
