package errcode

// 用户中心模块错误码, 模块代码为 01

const (
	// 用户
	IncorrectAccountOrPassword = 30101
	IncorrectOldPassword       = 30102
	IllegalUsername            = 30103
	ExistsUserName             = 30104
	ExistsUserEmail            = 30105
	ExistsUserPhone            = 30106
	// 文件
	FileUploadError           = 30107
	FileNotFound              = 30108
	FilePreviewParameterError = 30109
	FileNotSupportPreview     = 30110
)

const (
	IncorrectAccountOrPasswordMsg = "账号或密码错误"
	IncorrectOldPasswordMsg       = "旧密码错误"
	IllegalUsernameMsg            = "非法用户名 (%s)"
	ExistsUserNameMsg             = "用户名 (%s) 已被占用"
	ExistsUserEmailMsg            = "邮箱 (%s) 已被占用"
	ExistsUserPhoneMsg            = "手机号 (%s) 已被占用"
	FileUploadErrorMsg            = "文件上传失败，请稍后再试吧！"
	FileNotFoundMsg               = "该文件不存在"
	FilePreviewParameterErrorMsg  = "文件预览参数错误"
	FileNotSupportPreviewMsg      = "该文件不支持预览"
)
