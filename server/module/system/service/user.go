package service

import "github.com/antbiz/antadmin/module/system/service/dto"

// User 用户管理服务
var User = userService{}

type userService struct{}

// CreateUser 创建用户
func (s *userService) CreateUser(req *dto.User) (id string, err error) {
	return
}

// UpdateUser 更新用户
func (s *userService) UpdateUser(req *dto.User) (user *dto.User, err error) {
	return
}

// DeleteUser 删除用户
func (s *userService) DeleteUser(id string) error {
	return nil
}

// GetUser 用户详情
func (s *userService) GetUser(id string) (user *dto.User, err error) {
	return
}

// ListUser 用户列表
func (s *userService) ListUser() {

}
