package service

import (
	"context"

	"github.com/antbiz/antadmin/db"
	"github.com/antbiz/antadmin/db/umc/ent"
	"github.com/antbiz/antadmin/db/umc/ent/role"
	"github.com/antbiz/antadmin/module/umc/define"
)

// Role 角色管理服务
var Role = roleService{}

type roleService struct{}

// CreateRole 创建角色
func (s *roleService) CreateRole(ctx context.Context, req *define.CreateRoleRequest) (res *ent.Role, err error) {
	res, err = db.UmcClientSalve.Role.
		Create().
		SetName(req.Name).
		SetDisabled(req.Disabled).
		Save(ctx)
	return
}

// UpdateRole 更新角色
func (s *roleService) UpdateRole(ctx context.Context, id string, req *define.UpdateRoleRequest) (res *ent.Role, err error) {
	res, err = db.UmcClientSalve.Role.
		UpdateOneID(id).
		SetName(req.Name).
		SetDisabled(req.Disabled).
		Save(ctx)
	return
}

// DeleteRole 删除角色
func (s *roleService) DeleteRole(ctx context.Context, id string) error {
	err := db.UmcClientMaster.Role.
		DeleteOneID(id).
		Exec(ctx)
	return err
}

// GetRole 角色详情
func (s *roleService) GetRole(ctx context.Context, id string) (res *ent.Role, err error) {
	res, err = db.UmcClientSalve.Role.
		Get(ctx, id)
	return
}

// ListRole 用户列表
func (s *roleService) ListRole(ctx context.Context) (res []*ent.Role, qty int, err error) {
	res, err = db.UmcClientSalve.Role.
		Query().
		All(ctx)
	if err != nil {
		return
	}

	qty, err = db.UmcClientSalve.Role.
		Query().
		Count(ctx)
	if err != nil {
		return
	}

	return
}

// CheckRolename 检查角色名称是否符合规范(目前仅检查唯一性),存在返回false,否则true
func (s *roleService) CheckRolename(ctx context.Context, rolename string) bool {
	qty, err := db.UmcClientSalve.Role.
		Query().
		Where(
			role.NameEqualFold(rolename),
		).
		Count(ctx)
	if err != nil {
		return false
	}
	return qty == 0
}
