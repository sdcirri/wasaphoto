package database

func (db *appdbimpl) UsersExist(user1 string, user2 string) (bool, error) {
	exists1, err := db.UserExists(user1)
	if err != nil {
		return false, err
	}
	exists2, err := db.UserExists(user2)
	if err != nil {
		return false, err
	}
	return (exists1 && exists2), nil
}

