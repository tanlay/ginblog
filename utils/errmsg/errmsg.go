package errmsg

const (
	SUCCESS		= 000000
	ERROR 		= 999999
	//code=1000x,用户模块的错误
	ERROR_USER_EXIST		= 10001
	ERROR_USER_NOT_EXIST	= 10002
	ERROR_USERNAME_PASSWORD_WRONG	= 10003
	ERROR_USER_NO_PRIM	= 10004
	ERROR_TOKEN_EXIST 	= 10005
	ERROR_TOKEN_TIMEOUT 	= 10006
	ERROR_TOKEN_WRONG		= 10007
	ERROR_TOKEN_TYPE_ERONG	= 10008


	//code=2000x,文章模块的错误
	ERROR_ARTICLE_NOT_EXIST 	= 20003

	//code=3000x,分类模块的错误
	ERROR_CATEGORY_EXIST	= 30001
	ERROR_CATEGORY_NOT_EXIST	= 30002
)

var codeMsg = map[int]string {
	SUCCESS : "OK",
	ERROR: "FAIL",
	ERROR_USER_EXIST: "用户已存在",
	ERROR_USERNAME_PASSWORD_WRONG: "用户名密码错误",
	ERROR_USER_NOT_EXIST: "用户不存在",
	ERROR_USER_NO_PRIM: "用户无权限",
	ERROR_TOKEN_EXIST: "Token不存在",
	ERROR_TOKEN_TIMEOUT: "Token过期",
	ERROR_TOKEN_WRONG: "Token错误",
	ERROR_TOKEN_TYPE_ERONG: "Token格式错误",
	ERROR_CATEGORY_EXIST: "分类已存在",
	ERROR_CATEGORY_NOT_EXIST:"分类不存在",
	ERROR_ARTICLE_NOT_EXIST: "文章不存在",
}


func GetErrMsg(code int) string {
	return codeMsg[code]
}