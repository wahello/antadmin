package router

import (
	"github.com/antbiz/antadmin/middleware"
	"github.com/antbiz/antadmin/module/upms/api"
	"github.com/gogf/gf/net/ghttp"
)

// initUPMSRouter .
func initUPMSRouter(group *ghttp.RouterGroup) {
	group.POST("/api/user/signin", api.User.SignIn)

	group.Group("/api/upms", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)

		group.POST("/user", api.User.CreateUser)
		group.PUT("/user/:id", api.User.UpdateUser)
		group.DELETE("/user/:id", api.User.DeleteUser)
		group.GET("/user/:id", api.User.GetUser)
		group.GET("/user", api.User.ListUser)
		group.POST("/user/disabled", api.User.UpdateUserStatus)
	})
}
