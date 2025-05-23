package error

/*
 * basic code 100
 */

const (
	LoadConfigError      = 10001
	DBConnectError       = 10002 // 链接数据库错误
	AutoMigrateError     = 10003 // 自动迁移错误
	CreateNewUserError   = 10004 // 创建新用户错误
	GetUserError         = 10005 // 获取用户错误
	ChangePasswordError  = 10006 // 修改密码错误
	UpdateUserError      = 10007 // 更新用户错误
	DeleteUserError      = 10008 // 删除用户错误
	EncryptPasswordError = 10009 // 加密密码错误
	GenerateTokenError   = 10010 // 生成token错误
	ParseTokenError      = 10011 // 解析token错误
	RedisError           = 10012 // redis错误
	ReadFileError        = 10013 // 读取文件错误
	SendEmailError       = 10014 // 发送邮件错误
	MongoError           = 10015 // mongo错误
	ParseJsonError       = 10016 // 解析json错误
	RunCodeError         = 10017 // 运行代码错误
	CodeError            = 10018 // 代码错误
	WriteFileError       = 10019 // 写入文件错误
	CopyFileError        = 10020 // 复制文件错误
	StdinError           = 10021 // stdin错误
)
