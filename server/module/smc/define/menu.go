package define

// CreateMenuRequest 创建菜单请求参数
type CreateMenuRequest struct {
	Name   string `json:"name" v:"required#请输入菜单名称"`
	Hide   bool   `json:"hide"`
	Path   string `json:"path" v:"required#请输入菜单路由"`
	Sort   int    `json:"sort"`
	Icon   string `json:"icon"`
	Parent string `json:"parent"`
}

// UpdateMenuRequest 更新菜单请求参数
type UpdateMenuRequest struct {
	Name   string `json:"name" v:"required#请输入菜单名称"`
	Hide   bool   `json:"hide"`
	Path   string `json:"path" v:"required#请输入菜单路由"`
	Sort   int    `json:"sort"`
	Icon   string `json:"icon"`
	Parent string `json:"parent"`
}
