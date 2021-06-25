package errmsg

const (
	SUCCSE = 200
	ERROR  = 500

	// code=1XXX，用户模块错误
	ERROR_USERNAME_USED = 1001
	ERROR_PASSWORD_WRONG = 1002
	ERROR_USER_NOT_EXIST = 1003
	ERROR_TOKEN_EXIST = 1004
	ERROR_TOKEN_RUNTIME = 1005
	ERROR_TOKEN_WRONG = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007

	// code=2XXX，文章模块错误

	// code=3XXX，分类模块错误
)

var CodeMsg = map[int]string{
	SUCCSE:                 "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在！",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "TOKEN不存在,请重新登陆",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期,请重新登陆",
	ERROR_TOKEN_WRONG:      "TOKEN不正确,请重新登陆",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误,请重新登陆",
}

func GetErrorMsg(code int)string  {

	return CodeMsg[code]
}