package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// code= 1000... 用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	// code= 2000... 文章模块的错误

	// code= 3000... 分类模块的错误
	ERROR_CATENAME_USED = 2001
	// code = 9000... 通用错误
	ERROR_QUERY_WRONG    = 9001
	ERROR_VALIDATE_WRONG = 9002
)

var codeMsg = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",

	ERROR_USERNAME_USED:    "用户名已使用",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "token不存在",
	ERROR_TOKEN_RUNTIME:    "token已过期",
	ERROR_TOKEN_WRONG:      "token不正确",
	ERROR_TOKEN_TYPE_WRONG: "token格式错误",
	ERROR_CATENAME_USED:    "分类名已存在",
	ERROR_QUERY_WRONG:      "获取请求参数错误",
	ERROR_VALIDATE_WRONG:   "参数验证失败 ",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
