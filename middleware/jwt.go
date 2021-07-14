package middleware

import (
	"blog/model"
	"blog/utils"
	"blog/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)

var code int

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 生成token
func SetToken(username string) (string, int) {
	// 过期时间
	expirTime := time.Now().Add(10 * time.Hour)
	SetClaims := MyClaims{
		username,
		jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expirTime.Unix(),
			// 签发人
			Issuer: "Salmon",
		},
	}
	// 参数：签发方法，签发命令
	reClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	// 加签
	token,err := reClaim.SignedString(JwtKey)
	if err != nil{
		return "",errmsg.ERROR
	}
	return token,errmsg.SUCCSE
}

// 验证token
func CheckToken(token string) (*MyClaims, int) {
	settoken,_ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey,nil
	})
	if key,_ := settoken.Claims.(*MyClaims);settoken.Valid{
		return key,errmsg.SUCCSE
	}else {
		return nil,errmsg.ERROR
	}

}


// jwt中间件
func JwtToken() gin.HandlerFunc{
	return func(c *gin.Context) {
		// 接收参数
		tokenHeader := c.Request.Header.Get("Authorization")
		// 如果不存在
		if tokenHeader == ""{
			code = errmsg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code":code,
				"msg":errmsg.GetErrorMsg(code),
			})
			c.Abort()
			return
		}
		// 分割，验证格式
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":code,
				"msg":errmsg.GetErrorMsg(code),
			})
			c.Abort()
			return
		}
		// 验证token
		key,Tcode := CheckToken(checkToken[1])
		if Tcode == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":code,
				"msg":errmsg.GetErrorMsg(code),
			})
			c.Abort()
			return
		}
		// 验证权限
		var user model.User
		model.Db.Where("username=?",key.Username).First(&user)
		if user.Role != 1{
			c.JSON(http.StatusOK, gin.H{
				"code":errmsg.ERROR_USER_NO_RIGHT,
				"msg":errmsg.GetErrorMsg(errmsg.ERROR_USER_NO_RIGHT),
			})
			c.Abort()
			return
		}
		// 过期情况
		if time.Now().Unix() > key.ExpiresAt{
			code = errmsg.ERROR_TOKEN_RUNTIME
			c.JSON(http.StatusOK, gin.H{
				"code":code,
				"msg":errmsg.GetErrorMsg(code),
			})
			c.Abort()
			return
		}
		c.Set("username", key.Username)
		c.Next()
	}
}