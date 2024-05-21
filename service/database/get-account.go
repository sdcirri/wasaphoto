package database

import "database/sql"

func (db *appdbimpl) GetAccount(id string, username string) (Account, error) {
	var a Account
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

	err = db.c.QueryRow("select * from Users where username = ?", username).Scan(&a.Username, &a.ProPicPath)
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

	return a, nil
}
