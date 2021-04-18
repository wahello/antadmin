package errcode

// 常规的错误码, 模块代码为 00

const (
	// 系统内部错误代码
	UnknownError       = 10000
	ServerError        = 10001
	ServiceUnavailable = 10002
	TooManyRequests    = 10003
	CallHTTPError      = 10004
	DatabaseError      = 10005
	DaoCreateError     = 10006
	DaoUpdateError     = 10007
	DaoDeleteError     = 10008
	DaoGetError        = 10009
	DaoListError       = 10010

	// 系统外部错误代码
	IllegalRequest         = 20005
	AuthorizationError     = 20006
	IPDenied               = 20007
	PermissionDenied       = 20008
	IPRequestsLimit        = 20009
	UserRequestsLimit      = 20010
	APINotFound            = 20011
	RequestMethodError     = 20012
	ParameterBindError     = 20013
	MissRequiredParameter  = 20014
	ResubmitError          = 20015
	JSONError              = 20016
	DuplicateError         = 20017
	SourceNotFound         = 20018
	TimestampMismatchError = 20019
)

const (
	UnknownErrorMsg       = "未知错误"
	ServerErrorMsg        = "Oops，服务器居然开小差了，请稍后再试吧！"
	ServiceUnavailableMsg = "服务不可用"
	TooManyRequestsMsg    = "当前请求过多，系统繁忙"
	CallHTTPErrorMsg      = "调用第三方HTTP接口失败"
	DatabaseErrorMsg      = "数据库错误"
	DaoCreateErrorMsg     = "创建失败"
	DaoUpdateErrorMsg     = "更新失败"
	DaoDeleteErrorMsg     = "删除失败"
	DaoGetErrorMsg        = "查询详情失败"
	DaoListErrorMsg       = "查询列表失败"

	IllegalRequestMsg         = "非法请求"
	AuthorizationErrorMsg     = "未登录或非法访问"
	IPDeniedMsg               = "IP(%s) 限制不能请求该资源"
	PermissionDeniedMsg       = "权限不足"
	IPRequestsLimitMsg        = "IP(%s) 请求频次超过上限"
	UserRequestsLimitMsg      = "用户 (%s) 请求频次超过上限"
	APINotFoundMsg            = "你迷路了，当前请求: %s"
	RequestMethodErrorMsg     = "请求方法错误"
	ParameterBindErrorMsg     = "参数值非法，需为 (%s)，实际为 (%s)"
	MissRequiredParameterMsg  = "缺失必选参数 (%s)"
	ResubmitErrorMsg          = "请勿重复提交"
	JSONErrorMsg              = "无效的JSON"
	DuplicateErrorMsg         = "%s (%s) 已被占用"
	SourceNotFoundMsg         = "该资源不存在或权限不足"
	TimestampMismatchErrorMsg = "提交失败，该资源已被她/他人更新，请刷新后重试"
)
