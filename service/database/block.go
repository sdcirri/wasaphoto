package database

func (db *appdbimpl) Block(user int64, toBlock int64) error {
	exist, err := db.UsersExist(user, toBlock)
	if err != nil {
		return err
	}
	if !exist {
		return ErrUserNotFound
	}
	blocked, err := db.IsBlockedBy(toBlock, user)
	if err != nil {
		return err
	}
	if blocked {
		return ErrAlreadyBlocked
	}
	rmf_and_block, err := db.c.Begin()
	if err != nil {
		return err
	}
	_, err = rmf_and_block.Exec("insert into Blocks values (?, ?)", user, toBlock)
	if err != nil {
		err2 := rmf_and_block.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}
	_, err = rmf_and_block.Exec("delete from Follows where following = ? and follower = ?", user, toBlock)
	if err != nil {
		err2 := rmf_and_block.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}
	err = rmf_and_block.Commit()
	return err
}
