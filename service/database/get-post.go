package database

import "database/sql"

func (db *appdbimpl) GetPost(id string, postid int64) (Post, error) {
	var p Post
	err := db.c.QueryRow("select * from Posts where postID = ?", postid).Scan(&p.PostID, &p.ImgPath, &p.PubTime, &p.Author, &p.Caption)
	if err == sql.ErrNoRows {
		return p, ErrPostNotFound
	} else if err != nil {
		return p, err
	}
	blocked, err := db.IsBlockedBy(id, p.Author)
	if err != nil {
		// ErrUserNotFound: people not subscribed should not be able to view posts, so we treat it like an error
		p.ImgPath = ""			// Best to hide them even if p won't make it in the API response
		p.Caption = ""
		return p, err
	}
	if blocked {
		p.ImgPath = ""			// Best to hide them even if p won't make it in the API response
		p.Caption = ""
		return p, ErrUserIsBlocked
	}

	q, err := db.c.Query("select user from LikesP where post = ?", postid)
	if err != nil && err != sql.ErrNoRows {
		return p, err
	}
	p.Likes = make([]string, 0)
	for q.Next() {
		var like string
		err = q.Scan(&like)
		if err != nil {
			return p, err
		}
		p.Likes = append(p.Likes, like)
	}

	q, err = db.c.Query("select commentID from Comments where post = ?", postid)
	if err != nil && err != sql.ErrNoRows {
		return p, err
	}
	p.Comments = make([]int64, 0)
	for q.Next() {
		var com int64
		err = q.Scan(&com)
		if err != nil {
			return p, err
		}
		p.Comments = append(p.Comments, com)
	}

	return p, nil
}
