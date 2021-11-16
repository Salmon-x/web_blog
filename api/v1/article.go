package v1

import (
	"blog/model"
	"blog/utils"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ArticleApi struct {

}

// 添加文章
func (a *ArticleApi)AddArticle(c *gin.Context)  {
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
func (a *ArticleApi)GetArticles(c *gin.Context)  {
	// 字符串转int
	Size,_ := strconv.Atoi(c.DefaultQuery("size","3"))
	Page,_ := strconv.Atoi(c.DefaultQuery("page","1"))
	title := c.Query("title")
	if len(title) == 0 {
		data, code, total := model.GetArticles(Size, Page)
		c.JSON(http.StatusOK, gin.H{
			"code":  code,
			"msg":   errmsg.GetErrorMsg(code),
			"data":  data,
			"total": total,
		})
		return
	}
	data, code, total := model.SearchArticle(title,Size,Page)
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"data":    data,
		"total":   total,
		"msg": errmsg.GetErrorMsg(code),
	})

}

// 单个文章
func (a *ArticleApi)GetArticleInfo(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	data,code := model.ArticleInfo(id)
	// 渲染成html
	//data.Content = utils.Render(data.Content)

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":errmsg.GetErrorMsg(code),
		"data":data,
	})
}

func (a *ArticleApi)FrontArticleInfo(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	data,code := model.ArticleInfo(id)
	// 渲染成html
	data.Content = utils.Render(data.Content)

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":errmsg.GetErrorMsg(code),
		"data":data,
	})
}


// 分类下所有文章
func (a *ArticleApi)GetCateArt(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	Size,_ := strconv.Atoi(c.DefaultQuery("size","3"))
	Page,_ := strconv.Atoi(c.DefaultQuery("page","1"))
	data, code, total := model.GetCateArt(id, Size, Page)
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":errmsg.GetErrorMsg(code),
		"data":data,
		"total":total,
	})
}

// 文章编辑
func (a *ArticleApi)EditArticle(c *gin.Context)  {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code = model.UpdateArticle(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg": errmsg.GetErrorMsg(code),
	})
}

// 文章删除
func (a *ArticleApi)DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteArticle(id)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrorMsg(code),
	})
}
