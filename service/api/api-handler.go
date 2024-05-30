package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/login", rt.login)
	rt.router.PUT("/:userID/setPP", rt.setProPic)
	rt.router.GET("/users/:username", rt.getProfile)
	rt.router.PUT("/:userID/follow/:username", rt.follow)
	rt.router.DELETE("/:userID/unfollow/:username", rt.unfollow)
	rt.router.PUT("/:userID/block/:username", rt.block)
	rt.router.DELETE("/:userID/unblock/:username", rt.unblock)
	rt.router.GET("/:userID/followers", rt.getFollowers)
	rt.router.DELETE("/:userID/followers/:username/remove", rt.rmFollower)
	rt.router.POST("/:userID/newpost", rt.newPost)
	rt.router.GET("/feed", rt.getFeed)
	rt.router.GET("/posts/:postID", rt.getPost)
	rt.router.DELETE("/posts/:postID/delete", rt.rmPost)
	rt.router.GET("/posts/:postID/likes", rt.getLikes)
	rt.router.PUT("/posts/:postID/like", rt.likePost)
	rt.router.DELETE("/posts/:postID/unlike", rt.unlikePost)
	rt.router.POST("/posts/:postID/comment", rt.commentPost)
	rt.router.GET("/comments/:commentID", rt.getComment)
	rt.router.PUT("/comments/:commentID/like", rt.likeComment)
	rt.router.DELETE("/comments/:commentID/unlike", rt.unlikeComment)

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
