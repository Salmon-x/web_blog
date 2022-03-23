package v1

import (
	"blog/db"
	"blog/middleware"
	"blog/model"
	"blog/model/response"
	"blog/model/schemas"
	"blog/utils"
	"blog/utils/errmsg"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
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
			if utils.UseMultipoint{
				if err := db.RedisClient.Set(context.Background(),"admin", token, 10 * time.Hour).Err();err != nil {
					log.Println("设置状态错误: ",err)
					utils.Logger(err)
				}
			}
		}
		response.Result(code, errmsg.GetErrorMsg(code), token, c)
		return
	}
	response.Result(errmsg.ERROR_CAPTCHA_WRONG, errmsg.GetErrorMsg(errmsg.ERROR_CAPTCHA_WRONG), "", c)
}


func (l *LoginApi)LoginFront(c *gin.Context) {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)

	formData, code = model.CheckLoginFront(formData.Username, formData.Password)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    formData.Username,
		"id":      formData.ID,
		"message": errmsg.GetErrorMsg(code),
	})
}
