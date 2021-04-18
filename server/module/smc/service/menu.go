package service

import (
	"context"

	"github.com/antbiz/antadmin/db"
	"github.com/antbiz/antadmin/db/smc/ent"
	"github.com/antbiz/antadmin/module/smc/define"
)

// Menu 菜单管理服务
var Menu = menuService{}

type menuService struct{}

// CreateMenu 创建菜单
func (s *menuService) CreateMenu(ctx context.Context, req *define.CreateMenuRequest) (res *ent.Menu, err error) {
	res, err = db.SmcClientMaster.Menu.
		Create().
		SetName(req.Name).
		SetHide(req.Hide).
		SetPath(req.Path).
		SetSort(req.Sort).
		SetIcon(req.Icon).
		SetParent(req.Parent).
		Save(ctx)
	return
}

// UpdateMenu 更新菜单
func (s *menuService) UpdateMenu(ctx context.Context, id string, req *define.UpdateMenuRequest) (res *ent.Menu, err error) {
	res, err = db.SmcClientMaster.Menu.
		UpdateOneID(id).
		SetName(req.Name).
		SetHide(req.Hide).
		SetPath(req.Path).
		SetSort(req.Sort).
		SetIcon(req.Icon).
		SetParent(req.Parent).
		Save(ctx)
	return
}

// DeleteMenu 删除菜单
func (s *menuService) DeleteMenu(ctx context.Context, id string) error {
	err := db.SmcClientMaster.Menu.
		DeleteOneID(id).
		Exec(ctx)
	return err
}

// GetMenu 菜单详情
func (s *menuService) GetMenu(ctx context.Context, id string) (res *ent.Menu, err error) {
	res, err = db.SmcClientMaster.Menu.
		Get(ctx, id)
	return
}

// ListMenu 菜单列表
func (s *menuService) ListMenu(ctx context.Context) (res []*ent.Menu, qty int, err error) {
	res, err = db.SmcClientMaster.Menu.
		Query().
		All(ctx)
	if err != nil {
		return
	}

	qty, err = db.SmcClientMaster.Menu.
		Query().
		Count(ctx)
	if err != nil {
		return
	}

	return
}
