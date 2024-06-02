package database

func (db *appdbimpl) SearchUser(query string) ([]string, error) {
	res := make([]string, 0)
	q, err := db.c.Query("select username from Users where username like ? || '%'", query)
	if err != nil {
		return res, err
	}
	for q.Next() {
		var u string
		err = q.Scan(&u)
		if err != nil {
			return res, err
		}
		res = append(res, u)
	}
	return res, nil
}
