package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/session", rt.login)
	rt.router.PUT("/setUsername/:userID", rt.setUsername)
	rt.router.PUT("/setPP/:userID", rt.setProPic)
	rt.router.GET("/users/:userID", rt.getProfile)
	rt.router.POST("/users/:userID/follow/:toFollowID", rt.follow)
	rt.router.DELETE("/users/:userID/unfollow/:toUnfollowID", rt.unfollow)
	rt.router.POST("/users/:userID/block/:toBlockID", rt.block)
	rt.router.DELETE("/users/:userID/unblock/:toUnblockID", rt.unblock)
	rt.router.GET("/users/:userID/followers", rt.getFollowers)
	rt.router.GET("/users/:userID/following", rt.getFollowing)
	rt.router.GET("/users/:userID/blocked", rt.getBlocked)
	rt.router.DELETE("/users/:userID/followers/:toRemoveID/remove", rt.rmFollower)
	rt.router.POST("/users/:userID/newpost", rt.newPost)
	rt.router.GET("/posts/:postID", rt.getPost)
	rt.router.DELETE("/posts/:postID/delete", rt.rmPost)
	rt.router.GET("/posts/:postID/likes", rt.getLikes)
	rt.router.GET("/feed/:userID", rt.getFeed)
	rt.router.GET("/posts/:postID/liked/:userID", rt.isLiked)
	rt.router.PUT("/posts/:postID/like/:userID", rt.likePost)
	rt.router.DELETE("/posts/:postID/unlike/:userID", rt.unlikePost)
	rt.router.POST("/posts/:postID/comment/:userID", rt.commentPost)
	rt.router.GET("/comments/:commentID", rt.getComment)
	rt.router.PUT("/comments/:commentID/like/:userID", rt.likeComment)
	rt.router.DELETE("/comments/:commentID/unlike/:userID", rt.unlikeComment)
	rt.router.GET("/comments/:commentID/liked/:userID", rt.isCommentLiked)
	rt.router.DELETE("/comments/:commentID/delete/:userID", rt.deleteComment)
	rt.router.GET("/searchUser", rt.SearchUser)

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
