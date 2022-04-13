package v1

import (
	"blog/model/response"
	"blog/model/schemas"
	"blog/utils"
	"blog/utils/errmsg"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

// 存储store
var store = base64Captcha.DefaultMemStore

//var store = captcha.NewDefaultRedisStore()

type CaptCapApi struct {
}

func (u *CaptCapApi) GetCaptCha(c *gin.Context) {
	// 新的图片
	driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	//cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))  // 使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		utils.Logger(err)
		response.Result(errmsg.ERROR, "", c)
	} else {
		response.Result(errmsg.SUCCSE, response.SysCaptchaResponse{
			CaptchaId: id,
			PicPath:   b64s,
		}, c)
	}
}

func (u *CaptCapApi) VerifyCaptCha(c *gin.Context) {
	var data schemas.CaptChaModel
	_ = c.ShouldBindJSON(&data)
	if store.Verify(data.CaptchaId, data.Captcha, true) {
		response.Result(errmsg.SUCCSE, "", c)
		return
	}
	response.Result(errmsg.ERROR_CAPTCHA_WRONG, "", c)
}
