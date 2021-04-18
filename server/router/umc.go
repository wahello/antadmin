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
		group.GET("/profile", api.User.GetProfile)
		group.PUT("/profile", api.User.UpdateProfile)
		group.PUT("/update_password", api.User.UpdatePassword)
	})

	group.Group("/api/umc", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)

		// 用户模块
		group.POST("/user", api.User.CreateUser)
		group.PUT("/user/:id", api.User.UpdateUser)
		group.DELETE("/user/:id", api.User.DeleteUser)
		group.GET("/user/:id", api.User.GetUser)
		group.GET("/user", api.User.ListUser)

		// 角色模块
		group.POST("/role", api.Role.CreateRole)
		group.PUT("/role/:id", api.Role.UpdateRole)
		group.DELETE("/role/:id", api.Role.DeleteRole)
		group.GET("/role/:id", api.Role.GetRole)
		group.GET("/role", api.Role.ListRole)
	})
}
