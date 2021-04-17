package service

import (
	"context"

	"github.com/antbiz/antadmin/db"
	"github.com/antbiz/antadmin/db/umc/ent"
	"github.com/antbiz/antadmin/db/umc/ent/user"
	"github.com/antbiz/antadmin/module/umc/define"
	"github.com/gogf/gf/crypto/gmd5"
)

// User 用户管理服务
var User = userService{}

type userService struct{}

// EncryptPassword 明文密码加密
func (s *userService) EncryptPassword(username, password string) string {
	return gmd5.MustEncrypt(username + password)
}

// CreateUser 创建用户
func (s *userService) CreateUser(ctx context.Context, req *define.CreateUserRequest) (res *ent.User, err error) {
	res, err = db.UmcClientMaster.User.
		Create().
		SetUsername(req.Username).
		SetPassword(s.EncryptPassword(req.Username, req.Password)).
		SetNillablePhone(req.Phone).
		SetNillableEmail(req.Email).
		SetAvatar(req.Avatar).
		SetGender(req.Gender).
		Save(ctx)
	return
}

// UpdateUser 更新用户
func (s *userService) UpdateUser(ctx context.Context, req *define.UpdateUserRequest) (res *ent.User, err error) {
	res, err = db.UmcClientMaster.User.
		UpdateOneID(req.ID).
		SetNillablePhone(req.Phone).
		SetNillableEmail(req.Email).
		SetAvatar(req.Avatar).
		SetGender(req.Gender).
		Save(ctx)
	return
}

// DeleteUser 删除用户
func (s *userService) DeleteUser(ctx context.Context, id string) error {
	err := db.UmcClientMaster.User.
		DeleteOneID(id).
		Exec(ctx)
	return err
}

// GetUser 用户详情
func (s *userService) GetUser(ctx context.Context, id string) (res *ent.User, err error) {
	res, err = db.UmcClientSalve.User.
		Get(ctx, id)
	return
}

// ListUser 用户列表
func (s *userService) ListUser(ctx context.Context) (res []*ent.User, qty int, err error) {
	res, err = db.UmcClientSalve.User.
		Query().
		All(ctx)
	if err != nil {
		return
	}

	qty, err = db.UmcClientSalve.User.
		Query().
		Count(ctx)
	if err != nil {
		return
	}

	return
}

// FindUserByAccount 用户账号查询: 一般用于登录/注册
func (s *userService) FindUserByAccount(ctx context.Context, account string) (res *ent.User, err error) {
	res, err = db.UmcClientMaster.User.
		Query().
		Where(
			user.Or(
				user.Username(account),
				user.Phone(account),
				user.Email(account),
			),
		).
		Only(ctx)
	return
}

// CheckUsername 检查用户名是否符合规范(目前仅检查唯一性),存在返回false,否则true
func (s *userService) CheckUsername(ctx context.Context, username string) bool {
	qty, err := db.UmcClientSalve.User.
		Query().
		Where(
			user.UsernameEqualFold(username),
		).
		Count(ctx)
	if err != nil {
		return false
	}
	return qty == 0
}

// CheckPhone 检查手机号是否符合规范(目前仅检查唯一性),存在返回false,否则true
func (s *userService) CheckPhone(ctx context.Context, phone string) bool {
	qty, err := db.UmcClientSalve.User.
		Query().
		Where(
			user.PhoneEqualFold(phone),
		).
		Count(ctx)
	if err != nil {
		return false
	}
	return qty == 0
}

// CheckEmail 检查邮箱是否符合规范(目前仅检查唯一性),存在返回false,否则true
func (s *userService) CheckEmail(ctx context.Context, email string) bool {
	qty, err := db.UmcClientSalve.User.
		Query().
		Where(
			user.EmailEqualFold(email),
		).
		Count(ctx)
	if err != nil {
		return false
	}
	return qty == 0
}

// UpdateUserStatus 更新用户状态：启用/禁用
func (s *userService) UpdateUserStatus(ctx context.Context, req *define.UpdateUserStatusRequest) error {
	_, err := db.UmcClientMaster.User.
		UpdateOneID(req.ID).
		SetDisabled(req.Disabled).
		Save(ctx)
	return err
}
