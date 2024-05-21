package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
//	rt.router.GET("/", rt.feed)
//	rt.router.GET("/context", rt.wrap(rt.getContextReply))
	rt.router.GET("/login", rt.login)
	rt.router.POST("/register", rt.register)
	rt.router.POST("/setPP", rt.setProPic)
	rt.router.GET("/users/:username", rt.getProfile)
	rt.router.PUT("/users/:username/follow", rt.follow)
	rt.router.PUT("/users/:username/unfollow", rt.unfollow)
	rt.router.PUT("/users/:username/block", rt.block)
	rt.router.PUT("/users/:username/unblock", rt.unblock)
	rt.router.GET("/followers", rt.getFollowers)
	rt.router.PUT("/followers/:username/remove", rt.rmFollower)

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
