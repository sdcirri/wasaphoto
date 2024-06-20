package database

func (db *appdbimpl) IsCommentLiked(user int64, comment int64) (bool, error) {
	var ok bool
	err := db.c.QueryRow("select exists(select 1 from LikesC where user = ? and comment = ?)", user, comment).Scan(&ok)
	if err != nil {
		return false, err
	}
	return ok, nil
}
