package v1

import (
	"blog/model"
	"blog/utils/errmsg"
	"blog/utils/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


type UserApi struct {
}

var code int

// 添加用户
func (u *UserApi)AddUser(c *gin.Context)  {
	var data model.User
	var msg string
	_ = c.ShouldBindJSON(&data)
	msg, code = validator.Validate(&data)
	if code != errmsg.SUCCSE {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  code,
				"msg": msg,
			},
		)
		c.Abort()
		return
	}

	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCSE {
		model.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":errmsg.GetErrorMsg(code),
	})
}

// 查询用户
func (u *UserApi)GetUserInfo(c *gin.Context){
	id,_ := strconv.Atoi(c.Param("id"))
	data,code := model.GetUser(id)
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":errmsg.GetErrorMsg(code),
		"data":data,
	})
}

// 查询多个用户
func (u *UserApi)GetUsers(c *gin.Context)  {
	// 字符串转int
	Size,_ := strconv.Atoi(c.DefaultQuery("size","3"))
	Page,_ := strconv.Atoi(c.DefaultQuery("page","1"))
	username := c.Query("username")


	data,total := model.GetUsers(username,Size,Page)
	code = errmsg.SUCCSE
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":errmsg.GetErrorMsg(code),
		"data":data,
		"total":total,
	})


}

// 用户编辑
func (u *UserApi)EditUser(c *gin.Context)  {
	var data model.User
	id,_ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code = model.UniqueUser(data.Username, id)
	if code == errmsg.SUCCSE{
		model.UpdateUser(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED{
		c.Abort()
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":errmsg.GetErrorMsg(code),
	})
}

// 用户删除
func (u *UserApi)DeleteUser(c *gin.Context)  {
	// 接收id
	id,_ := strconv.Atoi(c.Param("id"))
	// 执行方法
	code = model.DeleteUser(id)
	// 返回
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":errmsg.GetErrorMsg(code),
	})
}

// 管理员设置密码
func (u *UserApi)AdminEditPass(c *gin.Context)  {
	var password map[string]string
	_ = c.ShouldBindJSON(&password)
	id,_ := strconv.Atoi(c.Param("id"))
	code = model.AdminEdit(id,password["password"])
	// 返回
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":errmsg.GetErrorMsg(code),
	})
}