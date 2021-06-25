package v1

import (
	"blog/model"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

// 添加用户
func AddUser(c *gin.Context)  {
	var data model.User
	_ = c.ShouldBindJSON(&data)
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
		"data":data,
	})
}

// 查询用户

// 查询多个用户
func GetUsers(c *gin.Context)  {
	// 字符串转int
	PageSize,_ := strconv.Atoi(c.DefaultQuery("pagesize","3"))
	PageNum,_ := strconv.Atoi(c.DefaultQuery("pagenum","1"))


	data := model.GetUsers(PageSize,PageNum)
	code = errmsg.SUCCSE
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":errmsg.GetErrorMsg(code),
		"data":data,
	})


}

// 用户编辑
func EditUser(c *gin.Context)  {
	
}

// 用户删除
func DeleteUser(c *gin.Context)  {
	
}