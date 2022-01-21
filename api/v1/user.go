package v1

import (
	"blog/model"
	"blog/model/response"
	"blog/utils/errmsg"
	"blog/utils/validator"
	"github.com/gin-gonic/gin"
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
		response.Result(code, msg, "", c)
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

	response.Result(code, msg, data, c)
}

// 查询用户
func (u *UserApi)GetUserInfo(c *gin.Context){
	id,_ := strconv.Atoi(c.Param("id"))
	data,code := model.GetUser(id)
	response.Result(code, errmsg.GetErrorMsg(code), data, c)
}

// 查询多个用户
func (u *UserApi)GetUsers(c *gin.Context)  {
	// 字符串转int
	Size,_ := strconv.Atoi(c.DefaultQuery("size","3"))
	Page,_ := strconv.Atoi(c.DefaultQuery("page","1"))
	username := c.Query("username")
	data,total := model.GetUsers(username,Size,Page)
	response.ResultAll(code, errmsg.GetErrorMsg(code), data, total, c)
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
	response.Result(code, errmsg.GetErrorMsg(code), "", c)
}

// 用户删除
func (u *UserApi)DeleteUser(c *gin.Context)  {
	// 接收id
	id,_ := strconv.Atoi(c.Param("id"))
	// 执行方法
	code = model.DeleteUser(id)
	// 返回
	response.Result(code, errmsg.GetErrorMsg(code), "", c)
}

// 管理员设置密码
func (u *UserApi)AdminEditPass(c *gin.Context)  {
	var password map[string]string
	_ = c.ShouldBindJSON(&password)
	id,_ := strconv.Atoi(c.Param("id"))
	code = model.AdminEdit(id,password["password"])
	// 返回
	response.Result(code, errmsg.GetErrorMsg(code), "", c)
}