package middleware

import (
	"github.com/antbiz/antadmin/common/errcode"
	"github.com/antbiz/antadmin/common/resp"
	"github.com/antbiz/antadmin/common/shared"
	"github.com/gogf/gf/net/ghttp"
)

// Auth 鉴权中间件，依赖上下文对象中间件
func Auth(r *ghttp.Request) {
	if customCtx := shared.Context.Get(r.Context()); customCtx == nil || customCtx.User == nil {
		resp.Error(r).SetCode(errcode.AuthorizationError).SetMsg(errcode.AuthorizationErrorMsg).Json()
	}
	r.Middleware.Next()
}
