package database

import (
	"encoding/base64"
	"database/sql"
	"io/ioutil"
)

func (db *appdbimpl) GetAccount(id string, username string) (Account, error) {
	var a Account
	var imgPath string
	exists, err := db.UserExists(username)
	if err != nil {
		return a, err
	}
	if exists {
		blocked, err := db.IsBlockedBy(id, username)
		if err != nil && err != ErrUserNotFound {
			return a, err
		}
		if blocked {
			return a, ErrUserIsBlocked
		}
	} else {
		return a, ErrUserNotFound
	}

	err = db.c.QueryRow("select username, propic from Users where username = ?", username).Scan(&a.Username, &imgPath)
	if err != nil {
		return a, err
	}
	err = db.c.QueryRow("select count(*) from Follows where following = ?", username).Scan(&a.Followers)
	if err == sql.ErrNoRows {
		a.Followers = 0
	} else if err != nil {
		return a, err
	}

	err = db.c.QueryRow("select count(*) from Follows where follower = ?", username).Scan(&a.Following)
	if err == sql.ErrNoRows {
		a.Following = 0
	} else if err != nil {
		return a, err
	}

	a.Posts = make([]int64, 0)
	q, err := db.c.Query("select postID from Posts where author = ?", username)
	if err != nil && err != sql.ErrNoRows {
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

	imgRaw, err := ioutil.ReadFile(imgPath)
	if err != nil {
		return a, err
	}
	a.ProPicB64 = base64.StdEncoding.EncodeToString(imgRaw)
	return a, nil
}
