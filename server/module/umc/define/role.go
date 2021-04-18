package define

// CreateRoleRequest 创建角色请求参数
type CreateRoleRequest struct {
	Name     string `json:"name" v:"required#请输入角色名称"`
	Disabled bool   `json:"disabled"`
}

// UpdateRoleRequest 更新角色请求参数
type UpdateRoleRequest struct {
	Name     string `json:"name" v:"required#请输入角色名称"`
	Disabled bool   `json:"disabled"`
}
