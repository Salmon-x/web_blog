package v1

import (
	"blog/middleware"
	"blog/model"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context)  {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)
	var token string
	var code int
	// 验证用户名密码
	formData, code = model.CheckLogin(formData.Username, formData.Password)
	if code == errmsg.SUCCSE{
		// 成功则签发token
		token, code = middleware.SetToken(formData.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMsg(code),
		"token": token,
	})
}