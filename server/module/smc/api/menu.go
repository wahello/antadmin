package api

import (
	"github.com/antbiz/antadmin/common/errcode"
	"github.com/antbiz/antadmin/common/resp"
	"github.com/antbiz/antadmin/module/smc/define"
	"github.com/antbiz/antadmin/module/smc/service"
	"github.com/gogf/gf/net/ghttp"
)

// Menu 菜单接口管理
var Menu = menuAPI{}

type menuAPI struct{}

// CreateMenu 创建菜单
func (a *menuAPI) CreateMenu(r *ghttp.Request) {
	var req *define.CreateMenuRequest
	if err := r.Parse(&req); err != nil {
		resp.Error(r).
			SetCode(errcode.ParameterBindError).
			SetError(err).
			Json()
	}

	if res, err := service.Menu.CreateMenu(r.Context(), req); err != nil {
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

// UpdateMenu 修改菜单
func (a *menuAPI) UpdateMenu(r *ghttp.Request) {
	var req *define.UpdateMenuRequest
	id := r.GetString("id")
	if err := r.Parse(&req); err != nil {
		resp.Error(r).
			SetCode(errcode.ParameterBindError).
			SetError(err).
			Json()
	}

	if res, err := service.Menu.UpdateMenu(r.Context(), id, req); err != nil {
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

// DeleteMenu 删除菜单
func (a *menuAPI) DeleteMenu(r *ghttp.Request) {
	id := r.GetString("id")
	if err := service.Menu.DeleteMenu(r.Context(), id); err != nil {
		resp.Error(r).
			SetCode(errcode.DaoDeleteError).
			SetMsg(errcode.DaoDeleteErrorMsg).
			Json()
	} else {
		resp.Success(r).Json()
	}
}

// GetMenu 菜单详情
func (a *menuAPI) GetMenu(r *ghttp.Request) {
	id := r.GetString("id")
	if res, err := service.Menu.GetMenu(r.Context(), id); err != nil {
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

// ListMenu 菜单列表
func (a *menuAPI) ListMenu(r *ghttp.Request) {
	if res, qty, err := service.Menu.ListMenu(r.Context()); err != nil {
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
