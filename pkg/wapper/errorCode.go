package wapper

type ErrorCode struct {
	Code    int
	Message string
}

var (
	//公共模块

	Success                = ErrorCode{Code: 10100, Message: "成功"}
	InternalServer         = ErrorCode{Code: 10001, Message: "系统内部错误"}
	ParameterMissing       = ErrorCode{Code: 10002, Message: "参数缺失"}
	ParameterFormat        = ErrorCode{Code: 10003, Message: "参数格式错误"}
	AuthenticationFailed   = ErrorCode{Code: 10004, Message: "认证失败（Token无效/过期）"}
	InsufficientAuthority  = ErrorCode{Code: 10005, Message: "权限不足"}
	DataNotFound           = ErrorCode{Code: 10006, Message: "数据未找到"}
	ParameterBindingFailed = ErrorCode{Code: 10007, Message: "参数绑定失败"}
	GetTotalRecordsFailed  = ErrorCode{Code: 10008, Message: "获取记录总数失败"}

	//注册模块
	RegisterFailed           = ErrorCode{Code: 20001, Message: "用户注册失败"}
	UsernameExisted          = ErrorCode{Code: 20002, Message: "用户已存在"}
	EmailExisted             = ErrorCode{Code: 20003, Message: "邮箱已存在"}
	PasswordEncryptionFailed = ErrorCode{Code: 20012, Message: "密码加密失败"}
	//登录模块
	UserNotFound          = ErrorCode{Code: 20004, Message: "用户不存在"}
	UsernameOrEmailIsNull = ErrorCode{Code: 200014, Message: "用户名或邮箱不能为空"}
	LoginFailed           = ErrorCode{Code: 20005, Message: "用户登录失败"}
	PasswordError         = ErrorCode{Code: 20006, Message: "密码错误"}
	NotBeenActivated      = ErrorCode{Code: 20007, Message: "用户未被激活"}
	//用户模块
	GetUserListFailed          = ErrorCode{Code: 20009, Message: "用户列表获取失败"}
	GetUserInfoFailed          = ErrorCode{Code: 200013, Message: "用户信息查询失败"}
	UpdateUserFailed           = ErrorCode{Code: 20010, Message: "用户信息更新失败"}
	UserInformationDiscrepancy = ErrorCode{Code: 20011, Message: "用户不一致"}
	DelUserFailed              = ErrorCode{Code: 20012, Message: "删除用户失败"}

	//资源模块
	GetResourceListFailed = ErrorCode{Code: 30001, Message: "资源列表获取失败"}
	AddResourceFailed     = ErrorCode{Code: 30002, Message: "资源创建失败"}
	GetResourceFailed     = ErrorCode{Code: 30004, Message: "资源列表获取失败"}
	UpdateResourceFailed  = ErrorCode{Code: 30003, Message: "资源更新失败"}
	DelResourceFailed     = ErrorCode{Code: 30005, Message: "资源删除失败"}

	//角色模块
	AddRoleFailed     = ErrorCode{Code: 40001, Message: "角色创建失败"}
	GetRoleFailed     = ErrorCode{Code: 40002, Message: "角色获取失败"}
	GetRoleListFailed = ErrorCode{Code: 40003, Message: "角色列表获取失败"}
	UpdateRoleFailed  = ErrorCode{Code: 40004, Message: "角色更新失败"}
	DelRoleFailed     = ErrorCode{Code: 400045, Message: "角色删除失败"}
	//角色-资源模块
	RoleResourceBindFailed = ErrorCode{Code: 50001, Message: "角色资源绑定失败"}
	AddRoleResourceFailed  = ErrorCode{Code: 50002, Message: "角色资源添加失败"}
	GetRoleResourceFailed  = ErrorCode{Code: 50003, Message: "查询角色拥有的资源失败"}
	//用户-角色模块
	AddUserRoleFailed = ErrorCode{Code: 60001, Message: "用户角色添加失败"}
	GetUserRoleFailed = ErrorCode{Code: 60002, Message: "查询用户拥有的角色失败"}

	//用户拥有的角色权限列表
	GrtUserResourceFailed = ErrorCode{Code: 70001, Message: "查询角色拥有的资源失败"}
)
