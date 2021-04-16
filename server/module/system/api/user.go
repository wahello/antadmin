package api

import (
	"github.com/antbiz/antadmin/common/errcode"
	"github.com/antbiz/antadmin/common/resp"
	"github.com/antbiz/antadmin/module/system/define"
	"github.com/antbiz/antadmin/module/system/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// User 用户接口管理
var User = userApi{}

type userApi struct{}

// @summary 用户创建接口
// @tags    用户服务
// @produce json
// @param   username formData string true "用户名"
// @param   password formData string true "明文密码"
// @param   phone formData string false "手机号"
// @param   email formData string false "邮箱"
// @param   avatar formData string false "头像"
// @param   gender formData int false "性别"
// @router  /user [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *userApi) CreateUser(r *ghttp.Request) {
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
			SetMsg(errcode.ExistsUserNameMsg).
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

// @summary 用户信息更新接口
// @tags    用户服务
// @produce json
// @param   phone formData string false "手机号"
// @param   email formData string false "邮箱"
// @param   avatar formData string false "头像"
// @param   gender formData int false "性别"
// @router  /user [PUT]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *userApi) UpdateUser(r *ghttp.Request) {
	var req *define.UpdateUserRequest
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

	if res, err := service.User.UpdateUser(r.Context(), req); err != nil {
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

// @summary 用户删除接口
// @tags    用户服务
// @produce json
// @router  /user/:id [DELETE]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *userApi) DeleteUser(r *ghttp.Request) {
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

// @summary 用户列表接口
// @tags    用户服务
// @produce json
// @router  /user [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *userApi) ListUser(r *ghttp.Request) {
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

// @summary 用户详情接口
// @tags    用户服务
// @produce json
// @router  /user/:id [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *userApi) GetUser(r *ghttp.Request) {
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

// @summary 用户状态更新接口
// @tags    用户服务
// @produce json
// @param   id formData string true "id"
// @param   disabled formData bool true "是否禁用"
// @router  /user/disabled [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *userApi) UpdateUserStatus(r *ghttp.Request) {
	var req *define.UpdateUserStatusRequest
	if err := r.Parse(&req); err != nil {
		resp.Error(r).
			SetCode(errcode.ParameterBindError).
			SetError(err).
			Json()
	}

	if err := service.User.UpdateUserStatus(r.Context(), req); err != nil {
		resp.Error(r).
			SetCode(errcode.DaoUpdateError).
			SetMsg(errcode.DaoUpdateErrorMsg).
			Json()
	} else {
		resp.Success(r).Json()
	}
}

// @summary 用户登录接口
// @tags    用户服务
// @produce json
// @param   login formData string true "用户名/手机号/邮箱"
// @param   password formData string true "用户密码"
// @router  /user/signin [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *userApi) SignIn(r *ghttp.Request) {
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
