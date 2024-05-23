/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
)

// Data model
type Account struct {
	Username	string		`json:"username"`
	ProPicB64	string		`json:"proPicB64"`
	Followers	uint		`json:"followers"`
	Following	uint		`json:"following"`
	Posts		[]int64		`json:"posts"`
}

type Post struct {
	PostID		int64		`json:"postID"`
	ImageB64	string		`json:"imageB64"`
	PubTime		string		`json:"pub_time"`
	Caption		string		`json:"caption"`
	Author		string		`json:"author"`
	Likes		[]string	`json:"likes"`
	Comments	[]int64		`json:"comments"`
}

type Comment struct {
	CommentID	int64		`json:"commentID"`
	PostID		int64		`json:"postID"`
	Author		string		`json:"author"`
	Time		string		`json:"time"`
	Content		string		`json:"content"`
	Likes		int64		`json:"likes"`
}

// Custom errors
var (
	ErrUserNotFound      = errors.New("Error: user does not exist")
	ErrAlreadyBlocked    = errors.New("Error: user is already blocked")
	ErrPostNotFound      = errors.New("Error: post does not exist")
	ErrUserIsBlocked     = errors.New("Error: user is blocked")
	ErrUserNotBlocked    = errors.New("Error: user is not blocked")
	ErrCommentNotFound   = errors.New("Error: comment not found")
	ErrUserAlreadyExists = errors.New("Error: user already exists")
	ErrNotFollowing      = errors.New("Cannot unfollow user: is not followed")
	ErrAlreadyFollowing  = errors.New("Error: already following")
	ErrDidNotLike        = errors.New("Error: user did not like post/comment")
	ErrBadImage          = errors.New("Error: bad image")
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	UserExists		(login string)								(bool, error)
	UsersExist		(user1 string, user2 string)				(bool, error)
	RegisterUser	(username string)							error
	SetProPic		(username string, imgpath string)			error
	Follows			(follower string, following string)			(bool, error)
	Follow			(follower string, toFollow string)			error
	Unfollow		(follower string, toUnfollow string)		error
	GetFollowers	(id string)									([]string, error)
	RmFollower		(user string, follower string)				error
	Block			(user string, toBlock string) 				error
	Unblock			(user string, toUnblock string) 			error
	IsBlockedBy		(blocked string, blocker string)			(bool, error)
	NewPost			(op string, imgpath string, caption string)	(int64, error)
	RmPost			(op string, postid int64)					error
	PostExists		(postID int64)								(bool, error)
	GetPost			(id string, postid int64)					(Post, error)
	GetAccount		(id string, username string)				(Account, error)
	CommentExists	(commentID int64)							(bool, error)
	GetComment		(commentID int64)							(Comment, error)
	LikePost		(user string, postID int64)					error
	UnlikePost		(user string, postID int64)					error
	CommentPost		(user string, postID int64, comment string)	(int64, error)
	LikeComment		(user string, commentID int64)				error
	UnlikeComment	(user string, commentID int64)				error
	GetFeed			(user string)								([]int64, error)
	Ping			()											error

}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// SQL statements for each table
	tables := [7]string{
`create table if not exists Users (
	username	varchar(64)		primary key,
	propic		varchar(255)	not null default '/srv/wasaphoto/propic_default.jpg'
);`,
`create table if not exists Follows (
	follower	varchar(64),
	following	varchar(64),
	foreign key (follower)  references User(login),
	foreign key (following) references User(login),
	primary key (follower, following),
	check (follower != following)
);`,
`create table if not exists Blocks (
	blocker varchar(64),
	blocked varchar(64),
	foreign key (blocker) references User(login),
	foreign key (blocked) references User(login),
	primary key (blocker, blocked),
	check (blocker != blocked)
);`,
`create table if not exists Posts (
	postID   integer      primary key autoincrement,
	img_path varchar(255) not null,
	pub_time datetime     not null,
	author   varchar(64)  not null,
	text     varchar(2048),
	foreign key (author) references User(login)
);`,
`create table if not exists Comments (
	commentID integer       primary key autoincrement,
	time      datetime      not null,
	author    varchar(64)   not null,
	post      integer       not null,
	text      varchar(2048) not null,
	foreign key (author) references User(login),
	foreign key (post) references Post(postID)
);`,
`create table if not exists LikesP (
	user varchar(64),
	post integer,
	foreign key (user) references User(login),
	foreign key (post) references Post(postID),
	primary key (user, post)
);`,
`create table if not exists LikesC (
	user    varchar(64),
	comment integer,
	foreign key (user) references User(login),
	foreign key (comment) references Comment(commentID),
	primary key (user, comment)
);`,
	}

	migration, err := db.Begin()
	if err != nil {
		return &appdbimpl{
			c: db,
		}, err
	}

	// Create the tables in order
	for _, s := range tables {
		_, err = migration.Exec(s)
		if err != nil {
			migration.Rollback()
			return &appdbimpl{
				c: db,
			}, err
		}
	}

	// Finally execute the transaction
	err = migration.Commit()
	if err != nil {
		migration.Rollback()
		return &appdbimpl{
			c: db,
		}, err
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
