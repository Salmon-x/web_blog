package schemas

type CaptChaModel struct {
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

type Login struct {
	CaptChaModel
	Username string `json:"username"`
	Password string `json:"password"`
}
