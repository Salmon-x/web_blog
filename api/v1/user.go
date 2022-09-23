package v1

import (
	"blog/model"
	"blog/model/response"
	"blog/utils/errmsg"
	"blog/utils/validator"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

// 添加用户
func (u *UserApi) AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	msg, code := validator.Validate(&data)
	if code != errmsg.SUCCSE {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"data": data,
			"msg":  msg,
		})
		c.Abort()
		return
	}
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCSE {
		data.CreateUser()
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}
	response.Result(code, data, c)
}

// 查询用户
func (u *UserApi) GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := model.User{}
	code := data.GetUser(id)
	response.Result(code, data, c)
}

// 查询多个用户
func (u *UserApi) GetUsers(c *gin.Context) {
	// 字符串转int
	Size, _ := strconv.Atoi(c.DefaultQuery("size", "3"))
	Page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	username := c.Query("username")
	data := make(model.Users, 0)
	total := data.GetUsers(username, Size, Page)
	response.ResultAll(errmsg.SUCCSE, data, total, c)
}

// 用户编辑
func (u *UserApi) EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := model.UniqueUser(data.Username, id)
	if code == errmsg.SUCCSE {
		data.UpdateUser(id)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	response.Result(code, "", c)
}

// 用户删除
func (u *UserApi) DeleteUser(c *gin.Context) {
	// 接收id
	id, _ := strconv.Atoi(c.Param("id"))
	// 执行方法
	code := model.DeleteUser(id)
	// 返回
	response.Result(code, "", c)
}

// 管理员设置密码
func (u *UserApi) AdminEditPass(c *gin.Context) {
	var password map[string]string
	_ = c.ShouldBindJSON(&password)
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.AdminEdit(id, password["password"])
	// 返回
	response.Result(code, "", c)
}
