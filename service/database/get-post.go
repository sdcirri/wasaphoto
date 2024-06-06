package database

import (
	"database/sql"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"time"
)

func (db *appdbimpl) GetPost(id int64, postid int64) (Post, error) {
	var p Post
	var imgPath string
	var pubts time.Time
	err := db.c.QueryRow("select * from Posts where postID = ?", postid).Scan(&p.PostID, &imgPath, &pubts, &p.Author, &p.Caption)
	if errors.Is(err, sql.ErrNoRows) {
		return p, ErrPostNotFound
	} else if err != nil {
		return p, err
	}
	blocked, err := db.IsBlockedBy(id, p.Author)
	if err != nil {
		return p, err
	}
	if blocked {
		return p, ErrUserIsBlocked
	}

	p.PubTime = pubts.Format(time.RFC3339)
	q, err := db.c.Query("select user from LikesP where post = ?", postid)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return p, err
	}
	p.Likes = make([]string, 0)
	for q.Next() {
		var like string
		err = q.Err()
		if err != nil {
			return p, err
		}
		err = q.Scan(&like)
		if err != nil {
			return p, err
		}
		p.Likes = append(p.Likes, like)
	}

	q, err = db.c.Query("select commentID from Comments where post = ?", postid)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return p, err
	}
	p.Comments = make([]int64, 0)
	for q.Next() {
		var com int64
		err = q.Err()
		if err != nil {
			return p, err
		}
		err = q.Scan(&com)
		if err != nil {
			return p, err
		}
		p.Comments = append(p.Comments, com)
	}

	imgRaw, err := ioutil.ReadFile(imgPath)
	if err != nil {
		return p, err
	}
	p.ImageB64 = base64.StdEncoding.EncodeToString(imgRaw)
	return p, nil
}
