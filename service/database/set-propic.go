package database

import "os"

func (db *appdbimpl) SetProPic(username string, imgpath string) error {
	_, err := os.Stat(imgpath)
	if err != nil {
		return err
	}
	exists, err := db.UserExists(username)
	if err != nil {
		return err
	}
	if !exists {
		return ErrUserNotFound
	}
	tran, err := db.c.Prepare("modify table Users set propic = ? where username = ?")
	if err != nil {
		return err
	}
	_, err = tran.Exec(imgpath, username)
	return err
}
