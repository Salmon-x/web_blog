package v1

import (
	"blog/middleware"
	"blog/model"
	"blog/model/response"
	"blog/model/schemas"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
)

type LoginApi struct {

}

func (l *LoginApi)Login(c *gin.Context)  {
	var (
		formData schemas.Login
		token string
		code int
		user model.User
	)
	_ = c.ShouldBindJSON(&formData)
	if store.Verify(formData.CaptchaId, formData.Captcha, true){
		// 验证用户名密码
		user, code = model.CheckLogin(formData)
		if code == errmsg.SUCCSE{
			// 成功则签发token
			token, code = middleware.SetToken(user.Username)
		}
		response.Result(code, errmsg.GetErrorMsg(code), token, c)
		return
	}
	response.Result(errmsg.ERROR_CAPTCHA_WRONG, errmsg.GetErrorMsg(errmsg.ERROR_CAPTCHA_WRONG), "", c)
}
