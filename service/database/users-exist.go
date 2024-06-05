package database

func (db *appdbimpl) UsersExist(user1 int64, user2 int64) (bool, error) {
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
