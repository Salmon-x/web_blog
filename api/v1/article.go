package v1

import (
	"blog/model"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加文章
func AddArticle(c *gin.Context)  {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	if code == errmsg.SUCCSE {
		model.CreateArticle(&data)
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

// 查询文章列表
func GetArticles(c *gin.Context)  {
	// 字符串转int
	PageSize,_ := strconv.Atoi(c.DefaultQuery("size","3"))
	PageNum,_ := strconv.Atoi(c.DefaultQuery("page","1"))


	data := model.GetArticles(PageSize,PageNum)
	code = errmsg.SUCCSE
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":errmsg.GetErrorMsg(code),
		"data":data,
	})


}

// 文章编辑
func EditArticle(c *gin.Context)  {
	var data model.Article
	id,_ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	if code == errmsg.SUCCSE{
		model.UpdateArticle(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED{
		c.Abort()
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":errmsg.GetErrorMsg(code),
	})
}

// 文章删除
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteArticle(id)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrorMsg(code),
	})
}