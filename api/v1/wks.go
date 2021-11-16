package v1

import (
	"blog/model"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type WksApi struct {

}

// 添加分类
func (w *WksApi)AddWks(c *gin.Context)  {
	var data model.WellKnownSaying
	_ = c.ShouldBindJSON(&data)
	code = model.Checkwks(data.Title)
	if code == errmsg.SUCCSE {
		model.Createwks(&data)
	}

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":errmsg.GetErrorMsg(code),
		"data":data,
	})
}


// 查询分类列表
func (w *WksApi)GetWks(c *gin.Context)  {
	data := model.GetWks()
	code = errmsg.SUCCSE
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":errmsg.GetErrorMsg(code),
		"data":data,
	})
}

// 查询单个分类信息
func (w *WksApi)GetWksInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := model.GetWksInfo(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrorMsg(code),
		},
	)
}

// 分类编辑
func (w *WksApi)EditWks(c *gin.Context)  {
	var data model.WellKnownSaying
	id,_ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code = model.Checkwks(data.Title)
	if code == errmsg.SUCCSE{
		model.UpdateWks(id, &data)
	}
	if code == errmsg.ERROR_WKS_USED{
		c.Abort()
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":errmsg.GetErrorMsg(code),
	})
}

// 分类删除
func (w *WksApi)DeleteWks(c *gin.Context)  {
	id,_ := strconv.Atoi(c.Param("id"))
	code = model.DeleteWks(id)

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":errmsg.GetErrorMsg(code),
	})
}