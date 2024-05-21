package database

import "github.com/sdgondola/wasaphoto/service/globaltime"

func (db *appdbimpl) CommentPost(user string, postID int64, comment string) (int64, error) {
	exists, err := db.UserExists(user)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, ErrUserNotFound
	}
	exists, err = db.PostExists(postID)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, ErrPostNotFound
	}
	ins, err := db.c.Prepare("insert into Comments values (?, ?, ?, ?) returning commentID")
	if err != nil {
		return 0, err
	}
	res, err := ins.Exec(globaltime.Now(), user, postID, comment)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
