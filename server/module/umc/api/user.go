package api

import (
	"github.com/antbiz/antadmin/common/errcode"
	"github.com/antbiz/antadmin/common/resp"
	"github.com/antbiz/antadmin/common/shared"
	"github.com/antbiz/antadmin/module/umc/define"
	"github.com/antbiz/antadmin/module/umc/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// User 用户接口管理
var User = userAPI{}

type userAPI struct{}

// checkSaveDuplicate 唯一性检查
func (a *userAPI) checkSaveDuplicate(r *ghttp.Request, username, email, phone string) {
	if username != "" && !service.User.CheckUsername(r.Context(), username) {
		resp.Error(r).
			SetCode(errcode.DuplicateError).
			SetMsgf(errcode.DuplicateErrorMsg, "用户名", username).
			Json()
	}

	if email != "" && !service.User.CheckEmail(r.Context(), email) {
		resp.Error(r).
			SetCode(errcode.DuplicateError).
			SetMsgf(errcode.DuplicateErrorMsg, "邮箱", email).
			Json()
	}

	if phone != "" && !service.User.CheckPhone(r.Context(), phone) {
		resp.Error(r).
			SetCode(errcode.DuplicateError).
			SetMsgf(errcode.DuplicateErrorMsg, "手机号", phone).
			Json()
	}
}

// CreateUser 创建用户
func (a *userAPI) CreateUser(r *ghttp.Request) {
	var req *define.CreateUserRequest
	if err := r.Parse(&req); err != nil {
		resp.Error(r).
			SetCode(errcode.ParameterBindError).
			SetError(err).
			Json()
	}

	a.checkSaveDuplicate(r, req.Username, *req.Email, *req.Phone)

	if res, err := service.User.CreateUser(r.Context(), req); err != nil {
		resp.Error(r).
			SetCode(errcode.DaoCreateError).
			SetMsg(errcode.DaoCreateErrorMsg).
			Json()
	} else {
		resp.Success(r).
			SetData(res).
			Json()
	}
}

// UpdateUser 修改用户信息
func (a *userAPI) UpdateUser(r *ghttp.Request) {
	var req *define.UpdateUserRequest
	id := r.GetString("id")
	if err := r.Parse(&req); err != nil {
		resp.Error(r).
			SetCode(errcode.ParameterBindError).
			SetError(err).
			Json()
	}

	a.checkSaveDuplicate(r, "", *req.Email, *req.Phone)

	if res, err := service.User.UpdateUser(r.Context(), id, req); err != nil {
		resp.Error(r).
			SetCode(errcode.DaoUpdateError).
			SetMsg(errcode.DaoUpdateErrorMsg).
			Json()
	} else {
		resp.Success(r).
			SetData(res).
			Json()
	}
}

// DeleteUser 删除用户
func (a *userAPI) DeleteUser(r *ghttp.Request) {
	id := r.GetString("id")
	if err := service.User.DeleteUser(r.Context(), id); err != nil {
		resp.Error(r).
			SetCode(errcode.DaoDeleteError).
			SetMsg(errcode.DaoDeleteErrorMsg).
			Json()
	} else {
		resp.Success(r).Json()
	}
}

// ListUser 用户列表
func (a *userAPI) ListUser(r *ghttp.Request) {
	if res, qty, err := service.User.ListUser(r.Context()); err != nil {
		resp.Error(r).
			SetCode(errcode.DaoListError).
			SetMsg(errcode.DaoListErrorMsg).
			Json()
	} else {
		resp.Success(r).
			SetData(
				resp.ListsData{
					List:  res,
					Total: qty,
				},
			).
			Json()
	}
}

// GetUser 用户详情
func (a *userAPI) GetUser(r *ghttp.Request) {
	id := r.GetString("id")
	if res, err := service.User.GetUser(r.Context(), id); err != nil {
		resp.Error(r).
			SetCode(errcode.DaoGetError).
			SetMsg(errcode.DaoGetErrorMsg).
			Json()
	} else {
		resp.Success(r).
			SetData(res).
			Json()
	}
}

// SignIn 用户登录
func (a *userAPI) SignIn(r *ghttp.Request) {
	var req *define.UserSignInRequest
	if err := r.Parse(&req); err != nil {
		resp.Error(r).
			SetCode(errcode.ParameterBindError).
			SetError(err).
			Json()
	}

	user, err := service.User.FindUserByAccount(r.Context(), req.Account)
	if err != nil {
		resp.Error(r).
			SetCode(errcode.IncorrectAccountOrPassword).
			SetMsg(errcode.IncorrectAccountOrPasswordMsg).
			Json()
	}
	if user.Password != service.User.EncryptPassword(user.Username, req.Password) {
		resp.Error(r).
			SetCode(errcode.IncorrectAccountOrPassword).
			SetMsg(errcode.IncorrectAccountOrPasswordMsg).
			Json()
	}

	sidInfo := g.Map{
		"id":       user.ID,
		"username": user.Username,
		"phone":    user.Phone,
		"email":    user.Email,
		"avatar":   user.Avatar,
		"disabled": user.Disabled,
		"sid":      r.Session.Id(),
	}
	r.Session.SetMap(sidInfo)
	resp.Success(r).
		SetData(sidInfo).
		Json()
}

// GetProfile 获取个人信息
func (a *userAPI) GetProfile(r *ghttp.Request) {
	if customCtx := shared.Context.Get(r.Context()); customCtx != nil && customCtx.User != nil {
		if res, err := service.User.GetUser(r.Context(), customCtx.User.ID); err != nil {
			resp.Error(r).
				SetCode(errcode.DaoGetError).
				SetMsg(errcode.DaoGetErrorMsg).
				Json()
		} else {
			resp.Success(r).
				SetData(res).
				Json()
		}
	} else {
		resp.Error(r).
			SetCode(errcode.AuthorizationError).
			SetMsg(errcode.AuthorizationErrorMsg).
			Json()
	}
}

// UpdateProfile 更新个人信息
func (a *userAPI) UpdateProfile(r *ghttp.Request) {
	if customCtx := shared.Context.Get(r.Context()); customCtx != nil && customCtx.User != nil {
		var req *define.UpdateProfileRequest
		if err := r.Parse(&req); err != nil {
			resp.Error(r).
				SetCode(errcode.ParameterBindError).
				SetError(err).
				Json()
		}

		a.checkSaveDuplicate(r, "", *req.Email, *req.Phone)

		if res, err := service.User.UpdateProfile(r.Context(), customCtx.User.ID, req); err != nil {
			resp.Error(r).
				SetCode(errcode.DaoUpdateError).
				SetMsg(errcode.DaoUpdateErrorMsg).
				Json()
		} else {
			resp.Success(r).
				SetData(res).
				Json()
		}
	} else {
		resp.Error(r).
			SetCode(errcode.AuthorizationError).
			SetMsg(errcode.AuthorizationErrorMsg).
			Json()
	}
}

// UpdatePassword 更新个人密码
func (a *userAPI) UpdatePassword(r *ghttp.Request) {
	if customCtx := shared.Context.Get(r.Context()); customCtx != nil && customCtx.User != nil {
		var req *define.UpdatePasswordRequest
		if err := r.Parse(&req); err != nil {
			resp.Error(r).
				SetCode(errcode.ParameterBindError).
				SetError(err).
				Json()
		}

		if u, err := service.User.GetUser(r.Context(), customCtx.User.ID); err != nil {
			resp.Error(r).
				SetCode(errcode.DaoGetError).
				SetMsg(errcode.DaoGetErrorMsg).
				Json()
		} else {
			if u.Password == service.User.EncryptPassword(u.Username, req.OldPassword) {
				newEncryptPassword := service.User.EncryptPassword(u.Username, req.NewPassword)
				if err := service.User.UpdateEncryptPassword(r.Context(), customCtx.User.ID, newEncryptPassword); err != nil {
					resp.Error(r).
						SetCode(errcode.DaoUpdateError).
						SetMsg(errcode.DaoUpdateErrorMsg).
						Json()
				} else {
					resp.Success(r).Json()
				}
			} else {
				resp.Error(r).
					SetCode(errcode.IncorrectOldPassword).
					SetMsg(errcode.IncorrectOldPasswordMsg).
					Json()
			}
		}
	} else {
		resp.Error(r).
			SetCode(errcode.AuthorizationError).
			SetMsg(errcode.AuthorizationErrorMsg).
			Json()
	}
}
