package define

// CreateUserRequest 创建用户请求参数
type CreateUserRequest struct {
	Username string   `json:"username" v:"required|length:2,16#请输入用户名|用户名长度应当在:min到:max之前"`
	Password string   `json:"password" v:"required|password#请输入密码|密码安全等级过低"`
	Phone    *string  `json:"phone,omitempty" v:"phone#手机号格式不正确"`
	Email    *string  `json:"email,omitempty" v:"email#邮箱格式不正确"`
	Avatar   string   `json:"avatar,omitempty"`
	Gender   int      `json:"gender" v:"between:0,2#性别参数大小为0到2"`
	Disabled bool     `json:"disabled"`
	Roles    []string `json:"roles"`
}

// UpdateUserRequest 更新用户请求参数
type UpdateUserRequest struct {
	Phone    *string  `json:"phone,omitempty" v:"phone#手机号格式不正确"`
	Email    *string  `json:"email,omitempty" v:"email#邮箱格式不正确"`
	Avatar   string   `json:"avatar,omitempty"`
	Gender   int      `json:"gender" v:"between:0,2#性别参数大小为0到2"`
	Disabled bool     `json:"disabled"`
	Roles    []string `json:"roles"`
}

// UpdateProfileRequest 用户更新个人信息请求参数
type UpdateProfileRequest struct {
	Phone  *string `json:"phone,omitempty" v:"phone#手机号格式不正确"`
	Email  *string `json:"email,omitempty" v:"email#邮箱格式不正确"`
	Avatar string  `json:"avatar,omitempty"`
	Gender int     `json:"gender" v:"between:0,2#性别参数大小为0到2"`
}

// UpdatePasswordRequest 用户更新个人密码请求参数
type UpdatePasswordRequest struct {
	OldPassword  string `json:"oldPassword" v:"required#请输入旧密码"`
	NewPassword  string `json:"newPassword" v:"required|password#请输入新密码|密码安全等级过低"`
	NewPassword2 string `json:"newPassword2" v:"required|same:newPassword#请二次输入新密码|两次密码输入不一致"`
}

// UserSignInRequest 用户登录请求参数
type UserSignInRequest struct {
	Account  string `json:"account" v:"required#请输入账号"`
	Password string `json:"password" v:"required#请输入密码"`
}
