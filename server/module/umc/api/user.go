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

// CreateUser 创建用户
func (userAPI) CreateUser(r *ghttp.Request) {
	var req *define.CreateUserRequest
	if err := r.Parse(&req); err != nil {
		resp.Error(r).
			SetCode(errcode.ParameterBindError).
			SetError(err).
			Json()
	}

	if !service.User.CheckUsername(r.Context(), req.Username) {
		resp.Error(r).
			SetCode(errcode.ExistsUserName).
			SetMsgf(errcode.ExistsUserNameMsg, req.Username).
			Json()
	}

	if req.Email != nil && !service.User.CheckEmail(r.Context(), *req.Email) {
		resp.Error(r).
			SetCode(errcode.ExistsUserEmail).
			SetMsgf(errcode.ExistsUserEmailMsg, *req.Email).
			Json()
	}

	if req.Phone != nil && !service.User.CheckPhone(r.Context(), *req.Phone) {
		resp.Error(r).
			SetCode(errcode.ExistsUserPhone).
			SetMsgf(errcode.ExistsUserPhoneMsg, *req.Phone).
			Json()
	}

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
func (userAPI) UpdateUser(r *ghttp.Request) {
	var req *define.UpdateUserRequest
	id := r.GetString("id")
	if err := r.Parse(&req); err != nil {
		resp.Error(r).
			SetCode(errcode.ParameterBindError).
			SetError(err).
			Json()
	}

	if req.Email != nil && !service.User.CheckEmail(r.Context(), *req.Email) {
		resp.Error(r).
			SetCode(errcode.ExistsUserEmail).
			SetMsg(errcode.ExistsUserEmailMsg).
			Json()
	}

	if req.Phone != nil && !service.User.CheckPhone(r.Context(), *req.Phone) {
		resp.Error(r).
			SetCode(errcode.ExistsUserPhone).
			SetMsg(errcode.ExistsUserPhoneMsg).
			Json()
	}

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
func (userAPI) DeleteUser(r *ghttp.Request) {
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
func (userAPI) ListUser(r *ghttp.Request) {
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
func (userAPI) GetUser(r *ghttp.Request) {
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
func (userAPI) SignIn(r *ghttp.Request) {
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
func (userAPI) GetProfile(r *ghttp.Request) {
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
func (userAPI) UpdateProfile(r *ghttp.Request) {
	if customCtx := shared.Context.Get(r.Context()); customCtx != nil && customCtx.User != nil {
		var req *define.UpdateProfileRequest
		if err := r.Parse(&req); err != nil {
			resp.Error(r).
				SetCode(errcode.ParameterBindError).
				SetError(err).
				Json()
		}

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
func (userAPI) UpdatePassword(r *ghttp.Request) {
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
