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
	"os"
)

// Data model
type Account struct {
	UserID    int64   `json:"userID"`
	Username  string  `json:"username"`
	ProPicB64 string  `json:"proPicB64"`
	Followers uint    `json:"followers"`
	Following uint    `json:"following"`
	Posts     []int64 `json:"posts"`
}

type Post struct {
	PostID   int64    `json:"postID"`
	ImageB64 string   `json:"imageB64"`
	PubTime  string   `json:"pubTime"`
	Caption  string   `json:"caption"`
	Author   int64    `json:"author"`
	Likes    []string `json:"likes"`
	Comments []int64  `json:"comments"`
}

type Comment struct {
	CommentID int64  `json:"commentID"`
	PostID    int64  `json:"postID"`
	Author    int64  `json:"author"`
	Time      string `json:"time"`
	Content   string `json:"content"`
	Likes     uint   `json:"likes"`
}

// Custom errors
var (
	ErrUserNotFound         = errors.New("error: user does not exist")
	ErrAlreadyBlocked       = errors.New("error: user is already blocked")
	ErrPostNotFound         = errors.New("error: post does not exist")
	ErrUserIsBlocked        = errors.New("error: user is blocked")
	ErrUserNotBlocked       = errors.New("error: user is not blocked")
	ErrCommentNotFound      = errors.New("error: comment not found")
	ErrUserAlreadyExists    = errors.New("error: user already exists")
	ErrNotFollowing         = errors.New("cannot unfollow user: is not followed")
	ErrAlreadyFollowing     = errors.New("error: already following")
	ErrDidNotLike           = errors.New("error: user did not like post/comment")
	ErrBadImage             = errors.New("error: bad image")
	ErrUserIsNotAuthor      = errors.New("error: you cannot delete somebody else's post")
	ErrAlreadyLiked         = errors.New("already liked")
	ErrBadCharset           = errors.New("bad charset")
	ErrUsernameAlreadyTaken = errors.New("username already taken")
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	UserExists(userID int64) (bool, error)
	UsersExist(user1 int64, user2 int64) (bool, error)
	UsernameTaken(login string) (bool, error)
	RegisterUser(username string) (int64, error)
	SetProPic(userID int64, imgB64 string) error
	SetUsername(userID int64, username string) error
	Follows(follower int64, following int64) (bool, error)
	Follow(follower int64, toFollow int64) error
	Unfollow(follower int64, toUnfollow int64) error
	GetFollowers(id int64) ([]int64, error)
	GetFollowing(id int64) ([]int64, error)
	RmFollower(user int64, follower int64) error
	Block(user int64, toBlock int64) error
	Unblock(user int64, toUnblock int64) error
	IsBlockedBy(blocked int64, blocker int64) (bool, error)
	NewPost(op int64, imgpath string, caption string) (int64, error)
	RmPost(op int64, postid int64) error
	PostExists(postID int64) (bool, error)
	GetPost(userID int64, postid int64) (Post, error)
	GetAccount(id int64, userID int64) (Account, error)
	CommentExists(commentID int64) (bool, error)
	GetComment(commentID int64) (Comment, error)
	IsLiked(user int64, post int64) (bool, error)
	LikePost(user int64, postID int64) error
	UnlikePost(user int64, postID int64) error
	CommentPost(user int64, postID int64, comment string) (int64, error)
	LikeComment(user int64, commentID int64) error
	UnlikeComment(user int64, commentID int64) error
	DeleteComment(user int64, commentID int64) error
	GetFeed(user int64) ([]int64, error)
	SearchUser(query string) ([]int64, error)
	Ping() error
}

type appdbimpl struct {
	c           *sql.DB
	installRoot string
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB, installRoot string) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}
	_, err := os.Stat(installRoot)
	if err != nil {
		return nil, errors.New("a valid installation path is required when building a database")
	}

	// SQL statements for each table
	tables := [7]string{
		`create table if not exists Users (
	userID		integer			primary key,
	username	varchar(40)		not null,
	propic		varchar(255)	not null,
	unique(username)
);`,
		`create table if not exists Follows (
	follower	integer,
	following	integer,
	foreign key (follower)  references User(userID),
	foreign key (following) references User(userID),
	primary key (follower, following),
	check (follower != following)
);`,
		`create table if not exists Blocks (
	blocker integer,
	blocked integer,
	foreign key (blocker) references User(userID),
	foreign key (blocked) references User(userID),
	primary key (blocker, blocked),
	check (blocker != blocked)
);`,
		`create table if not exists Posts (
	postID   integer      primary key autoincrement,
	img_path varchar(255) not null,
	pub_time datetime     not null,
	author   integer	  not null,
	text     varchar(2048),
	foreign key (author) references User(userID)
);`,
		`create table if not exists Comments (
	commentID integer       primary key autoincrement,
	time      datetime      not null,
	author    integer	    not null,
	post      integer       not null,
	text      varchar(2048) not null,
	foreign key (author) references User(userID),
	foreign key (post) references Post(postID)
);`,
		`create table if not exists LikesP (
	user integer,
	post integer,
	foreign key (user) references User(userID),
	foreign key (post) references Post(postID),
	primary key (user, post)
);`,
		`create table if not exists LikesC (
	user    integer,
	comment integer,
	foreign key (user) references User(userID),
	foreign key (comment) references Comment(commentID),
	primary key (user, comment)
);`,
	}

	migration, err := db.Begin()
	if err != nil {
		return nil, err
	}

	// Create the tables in order
	for _, s := range tables {
		_, err = migration.Exec(s)
		if err != nil {
			err2 := migration.Rollback()
			if err2 != nil {
				return nil, err2
			}
			return nil, err
		}
	}

	// Finally execute the transaction
	err = migration.Commit()
	if err != nil {
		err2 := migration.Rollback()
		if err2 != nil {
			return nil, err2
		}
		return nil, err
	}

	return &appdbimpl{
		c:           db,
		installRoot: installRoot,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
