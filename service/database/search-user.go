package database

func (db *appdbimpl) SearchUser(query string) ([]int64, error) {
	res := make([]int64, 0)
	q, err := db.c.Query("select userID from Users where username like ? || '%'", query)
	if err != nil {
		return res, err
	}
	for q.Next() {
		var u int64
		err = q.Scan(&u)
		if err != nil {
			return res, err
		}
		res = append(res, u)
	}
	return res, nil
}
