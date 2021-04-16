package router

import (
	"github.com/antbiz/antadmin/middleware"
	"github.com/antbiz/antadmin/module/system/api"
	"github.com/gogf/gf/net/ghttp"
)

// initSystemRouter 初始化系统模块路由
func initSystemRouter(group *ghttp.RouterGroup) {
	group.Group("/api/system", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)

		group.POST("/user", api.User.CreateUser)
		group.PUT("/user/:id", api.User.UpdateUser)
		group.DELETE("/user/:id", api.User.DeleteUser)
		group.GET("/user/:id", api.User.GetUser)
		group.GET("/user", api.User.ListUser)
		group.POST("/user/signin", api.User.SignIn)
	})
}
