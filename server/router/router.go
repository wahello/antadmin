package router

import (
	"github.com/antbiz/antadmin/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(
			middleware.Ctx,
			middleware.CORS,
			middleware.ErrorHandler,
		)

		// 初始化用户管理模块路由
		initUmcRouter(group)

		// 初始化系统管理模块路由
		initSmcRouter(group)
	})
}
