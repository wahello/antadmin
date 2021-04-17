package router

import (
	"github.com/antbiz/antadmin/middleware"
	"github.com/antbiz/antadmin/module/umc/api"
	"github.com/gogf/gf/net/ghttp"
)

// initUmcRouter .
func initUmcRouter(group *ghttp.RouterGroup) {
	group.Group("/api/user", func(group *ghttp.RouterGroup) {
		group.POST("/signin", api.User.SignIn)

		group.Middleware(middleware.Auth)
		group.GET("/profile", api.User.Profile)
	})

	group.Group("/api/umc", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)

		group.POST("/user", api.User.CreateUser)
		group.PUT("/user/:id", api.User.UpdateUser)
		group.DELETE("/user/:id", api.User.DeleteUser)
		group.GET("/user/:id", api.User.GetUser)
		group.GET("/user", api.User.ListUser)
		group.POST("/user/disabled", api.User.UpdateUserStatus)
	})
}
