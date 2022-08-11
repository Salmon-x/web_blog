package v1

import (
	"blog/model"
	"blog/model/response"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CategoryApi struct {
}

// 添加分类
func (a *CategoryApi) AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code := model.CheckCategory(data.Name)
	if code == errmsg.SUCCSE {
		data.CreateCategory()
	}
	if code == errmsg.ERROR_CATENAME_USED {
		code = errmsg.ERROR_CATENAME_USED
	}

	response.Result(code, data, c)
}

// 查询单个分类信息
func (a *CategoryApi) GetCateInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := model.GetCateInfo(id)

	response.Result(code, data, c)
}

// 查询分类列表
func (a *CategoryApi) GetCategorys(c *gin.Context) {
	// 字符串转int
	PageSize, _ := strconv.Atoi(c.DefaultQuery("size", "3"))
	PageNum, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	data, total := model.GetCategorys(PageSize, PageNum)
	response.ResultAll(errmsg.SUCCSE, data, total, c)
}

// 分类编辑
func (a *CategoryApi) EditCategory(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := model.CheckCategory(data.Name)
	if code == errmsg.SUCCSE {
		data.UpdateCategory(id)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	response.Result(code, "", c)
}

// 分类删除
func (a *CategoryApi) DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteCategory(id)
	response.Result(code, "", c)
}
