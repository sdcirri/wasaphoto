package database

func (db *appdbimpl) GetBlocked(id int64) ([]int64, error) {
	blocks := make([]int64, 0)
	exists, err := db.UserExists(id)
	if err != nil {
		return blocks, err
	}
	if !exists {
		return blocks, ErrUserNotFound
	}
	q, err := db.c.Query("select blocked from Blocks where blocker = ?", id)
	if err != nil {
		return blocks, err
	}
	for q.Next() {
		var f int64
		err = q.Err()
		if err != nil {
			return blocks, err
		}
		err = q.Scan(&f)
		if err != nil {
			return blocks, err
		}
		blocks = append(blocks, f)
	}
	return blocks, nil
}
