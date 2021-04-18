package router

import (
	"github.com/antbiz/antadmin/middleware"
	"github.com/antbiz/antadmin/module/smc/api"
	"github.com/gogf/gf/net/ghttp"
)

// initSmcRouter .
func initSmcRouter(group *ghttp.RouterGroup) {
	group.Group("/api/smc", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)

		// 菜单模块
		group.POST("/menu", api.Menu.CreateMenu)
		group.PUT("/menu/:id", api.Menu.UpdateMenu)
		group.DELETE("/menu/:id", api.Menu.DeleteMenu)
		group.GET("/menu/:id", api.Menu.GetMenu)
		group.GET("/menu", api.Menu.ListMenu)
	})
}
