package errmsg

const (
	SUCCSE = 200
	ERROR  = 500

	// code=1XXX，用户模块错误
	ERROR_USERNAME_USED = 1001
	ERROR_PASSWORD_WRONG = 1002
	ERROR_USER_NOT_EXIST = 1003
	ERROR_TOKEN_EXIST = 1004
	ERROR_TOKEN_BLACK = 1009
	ERROR_TOKEN_RUNTIME = 1005
	ERROR_TOKEN_WRONG = 1006
	ERROR_CAPTCHA_WRONG = 1010
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008

	// code=2XXX，文章模块错误
	ERROR_ART_NOT_EXIST = 2001
	// code= 3000... 分类模块的错误
	ERROR_CATENAME_USED  = 3001
	ERROR_CATE_NOT_EXIST = 3002

	// 名言错误
	ERROR_WKS_USED = 5002

	// 文件错误
	ERROR_FILE_EXIST = 4001
)

var CodeMsg = map[int]string{
	SUCCSE:                 "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在！",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_CAPTCHA_WRONG:   "验证码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_BLACK:   "非法TOKEN",
	ERROR_USER_NO_RIGHT:    "权限不足",
	ERROR_TOKEN_EXIST:      "TOKEN不存在,请重新登陆",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期,请重新登陆",
	ERROR_TOKEN_WRONG:      "TOKEN不正确,请重新登陆",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误,请重新登陆",
	ERROR_ART_NOT_EXIST: "文章不存在",
	ERROR_CATENAME_USED:  "该分类已存在",
	ERROR_CATE_NOT_EXIST: "该分类不存在",
	ERROR_WKS_USED: "该名言已存在",
	ERROR_FILE_EXIST: "文件不存在",
}

func GetErrorMsg(code int)string  {
	return CodeMsg[code]
}