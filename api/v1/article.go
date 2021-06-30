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
	model.CreateArticle(&data)
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}

	c.JSON(http.StatusOK,gin.H{
		"code":errmsg.SUCCSE,
		"msg":errmsg.GetErrorMsg(errmsg.SUCCSE),
		"data":data,
	})
}

// 查询文章列表
func GetArticles(c *gin.Context)  {
	// 字符串转int
	Size,_ := strconv.Atoi(c.DefaultQuery("size","3"))
	Page,_ := strconv.Atoi(c.DefaultQuery("page","1"))


	data,code := model.GetArticles(Size,Page)
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":errmsg.GetErrorMsg(code),
		"data":data,
	})
}

// 单个文章
func GetArticleInfo(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	data,code := model.ArticleInfo(id)
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":errmsg.GetErrorMsg(code),
		"data":data,
	})
}

// 分类下所有文章
func GetCateArt(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	Size,_ := strconv.Atoi(c.DefaultQuery("size","3"))
	Page,_ := strconv.Atoi(c.DefaultQuery("page","1"))
	data, code := model.GetCateArt(id, Size, Page)
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":errmsg.GetErrorMsg(code),
		"data":data,
	})
}

// 文章编辑
func EditArticle(c *gin.Context)  {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code = model.UpdateArticle(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMsg(code),
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