package database

func (db *appdbimpl) IsLiked(user int64, post int64) (bool, error) {
	var ok bool
	err := db.c.QueryRow("select exists(select 1 from LikesP where user = ? and post = ?)", user, post).Scan(&ok)
	if err != nil {
		return false, err
	}
	return ok, nil
}
