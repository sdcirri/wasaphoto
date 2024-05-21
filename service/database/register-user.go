package database

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
	return err
}
