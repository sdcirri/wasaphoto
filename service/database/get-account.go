package database

import (
	"database/sql"
	"encoding/base64"
	"errors"
	"os"
)

func (db *appdbimpl) GetAccount(id int64, userID int64) (Account, error) {
	var a Account
	var imgPath string
	exists, err := db.UserExists(userID)
	if err != nil {
		return a, err
	}
	if exists {
		blocked, err := db.IsBlockedBy(id, userID)
		if err != nil && !errors.Is(err, ErrUserNotFound) {
			return a, err
		}
		if blocked {
			return a, ErrUserIsBlocked
		}
	} else {
		return a, ErrUserNotFound
	}

	err = db.c.QueryRow("select userID, username, propic from Users where userID = ?", userID).Scan(&a.UserID, &a.Username, &imgPath)
	if err != nil {
		return a, err
	}
	err = db.c.QueryRow("select count(*) from Follows where following = ?", userID).Scan(&a.Followers)
	if errors.Is(err, sql.ErrNoRows) {
		a.Followers = 0
	} else if err != nil {
		return a, err
	}

	err = db.c.QueryRow("select count(*) from Follows where follower = ?", userID).Scan(&a.Following)
	if errors.Is(err, sql.ErrNoRows) {
		a.Following = 0
	} else if err != nil {
		return a, err
	}

	a.Posts = make([]int64, 0)
	q, err := db.c.Query("select postID from Posts where author = ?", userID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return a, err
	}
	for q.Next() {
		var post int64
		err = q.Scan(&post)
		if err != nil {
			return a, err
		}
		a.Posts = append(a.Posts, post)
	}

	imgRaw, err := os.ReadFile(imgPath)
	if err != nil {
		return a, err
	}
	a.ProPicB64 = base64.StdEncoding.EncodeToString(imgRaw)
	return a, nil
}
