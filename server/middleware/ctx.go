package middleware

import (
	"github.com/antbiz/antadmin/common/shared"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// Ctx 自定义上下文对象
func Ctx(r *ghttp.Request) {
	customCtx := &shared.Ctx{
		Session: r.Session,
		Data:    make(g.Map),
	}
	shared.Context.Init(r, customCtx)

	if sessionUserID := r.Session.GetString("id"); sessionUserID != "" {
		customCtx.User = &shared.CtxUser{
			ID:       sessionUserID,
			Username: r.Session.GetString("username"),
			Phone:    r.Session.GetString("phone"),
			Email:    r.Session.GetString("email"),
			Avatar:   r.Session.GetString("avatar"),
			Disabled: r.Session.GetBool("disabled"),
			IsAdmin:  r.Session.GetBool("isAdmin"),
		}
	}

	// 将自定义的上下文对象传递到模板变量中使用
	r.Assigns(g.Map{
		"Context": customCtx,
	})

	r.Middleware.Next()
}
