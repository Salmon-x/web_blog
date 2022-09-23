package v1

import (
	"blog/model"
	"blog/model/response"
	"blog/utils/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WksApi struct {
}

// 添加分类
func (w *WksApi) AddWks(c *gin.Context) {
	var data model.WellKnownSaying
	_ = c.ShouldBindJSON(&data)
	code := model.Checkwks(data.Title)
	if code == errmsg.SUCCSE {
		data.Createwks()
	}
	response.Result(code, data, c)
}

// 查询分类列表
func (w *WksApi) GetWks(c *gin.Context) {
	data := make(model.WellKnownSayings, 0)
	data.GetWks()
	code := errmsg.SUCCSE
	response.Result(code, data, c)
}

// 查询单个分类信息
func (w *WksApi) GetWksInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := model.WellKnownSaying{}
	code := data.GetWksInfo(id)
	response.Result(code, data, c)
}

// 分类编辑
func (w *WksApi) EditWks(c *gin.Context) {
	var data model.WellKnownSaying
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := model.Checkwks(data.Title)
	if code == errmsg.SUCCSE {
		data.UpdateWks(id)
	}
	if code == errmsg.ERROR_WKS_USED {
		c.Abort()
	}
	response.Result(code, "", c)
}

// 分类删除
func (w *WksApi) DeleteWks(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteWks(id)
	response.Result(code, "", c)
}
