package errmsg

const (
	SUCCESS		= 000000
	ERROR 		= 999999
	//code=1000x,用户模块的错误
	ERROR_USER_EXIST		= 10001
	ERROR_USERNAME_PASSWORD_WRONG	= 10002
	ERROR_USER_NOT_EXIST	= 10003
	ERROR_TOKEN_NOT_EXIST 	= 10004
	ERROR_TOKEN_TIMEOUT 	= 10005
	ERROR_TOKEN_WRONG		= 10006
	ERROR_TOKEN_TYPE_ERONG	= 10007

	//code=2000x,文章模块的错误
	//code=3000x,分类模块的错误
)

var codeMsg = map[int]string {
	SUCCESS : "OK",
	ERROR: "FAIL",
	ERROR_USER_EXIST: "用户已存在",
	ERROR_USERNAME_PASSWORD_WRONG: "用户名密码错误",
	ERROR_USER_NOT_EXIST: "用户不存在",
	ERROR_TOKEN_NOT_EXIST: "Token不存在",
	ERROR_TOKEN_TIMEOUT: "Token过期",
	ERROR_TOKEN_WRONG: "Token错误",
	ERROR_TOKEN_TYPE_ERONG: "Token格式错误",
}


func GetErrMsg(code int) string {
	return codeMsg[code]
}