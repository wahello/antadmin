package api

import (
	"github.com/antbiz/antadmin/common/errcode"
	"github.com/antbiz/antadmin/common/resp"
	"github.com/antbiz/antadmin/module/umc/define"
	"github.com/antbiz/antadmin/module/umc/service"
	"github.com/gogf/gf/net/ghttp"
)

// Role 角色接口管理
var Role = roleAPI{}

type roleAPI struct{}

// checkSaveDuplicate 唯一性检查
func (a *roleAPI) checkSaveDuplicate(r *ghttp.Request, rolename string) {
	if rolename != "" && !service.Role.CheckRolename(r.Context(), rolename) {
		resp.Error(r).
			SetCode(errcode.DuplicateError).
			SetMsgf(errcode.DuplicateErrorMsg, "角色名称", rolename).
			Json()
	}
}

// CreateRole 创建角色
func (a *roleAPI) CreateRole(r *ghttp.Request) {
	var req *define.CreateRoleRequest
	if err := r.Parse(&req); err != nil {
		resp.Error(r).
			SetCode(errcode.ParameterBindError).
			SetError(err).
			Json()
	}

	a.checkSaveDuplicate(r, req.Name)

	if res, err := service.Role.CreateRole(r.Context(), req); err != nil {
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

// UpdateRole 修改角色
func (a *roleAPI) UpdateRole(r *ghttp.Request) {
	var req *define.UpdateRoleRequest
	id := r.GetString("id")
	if err := r.Parse(&req); err != nil {
		resp.Error(r).
			SetCode(errcode.ParameterBindError).
			SetError(err).
			Json()
	}

	a.checkSaveDuplicate(r, req.Name)

	if res, err := service.Role.UpdateRole(r.Context(), id, req); err != nil {
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

// DeleteRole 删除角色
func (a *roleAPI) DeleteRole(r *ghttp.Request) {
	id := r.GetString("id")
	if err := service.Role.DeleteRole(r.Context(), id); err != nil {
		resp.Error(r).
			SetCode(errcode.DaoDeleteError).
			SetMsg(errcode.DaoDeleteErrorMsg).
			Json()
	} else {
		resp.Success(r).Json()
	}
}

// GetRole 角色详情
func (a *roleAPI) GetRole(r *ghttp.Request) {
	id := r.GetString("id")
	if res, err := service.Role.GetRole(r.Context(), id); err != nil {
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

// ListRole 角色列表
func (a *roleAPI) ListRole(r *ghttp.Request) {
	if res, qty, err := service.Role.ListRole(r.Context()); err != nil {
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
