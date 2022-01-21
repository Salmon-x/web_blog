package v1

import (
	"blog/model"
	"blog/model/response"
	"blog/utils"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
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

	response.Result(code, errmsg.GetErrorMsg(code), data, c)
}

// 查询文章列表
func (a *ArticleApi)GetArticles(c *gin.Context)  {
	// 字符串转int
	Size,_ := strconv.Atoi(c.DefaultQuery("size","3"))
	Page,_ := strconv.Atoi(c.DefaultQuery("page","1"))
	title := c.Query("title")
	if len(title) == 0 {
		data, code, total := model.GetArticles(Size, Page)
		response.ResultAll(code, errmsg.GetErrorMsg(code), data, total, c)
		return
	}
	data, code, total := model.SearchArticle(title,Size,Page)
	response.ResultAll(code, errmsg.GetErrorMsg(code), data, total, c)
}

// 单个文章
func (a *ArticleApi)GetArticleInfo(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	data,code := model.ArticleInfo(id)
	// 渲染成html
	//data.Content = utils.Render(data.Content)

	response.Result(code, errmsg.GetErrorMsg(code), data, c)
}

func (a *ArticleApi)FrontArticleInfo(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	data,code := model.ArticleInfo(id)
	// 渲染成html
	data.Content = utils.Render(data.Content)

	response.Result(code, errmsg.GetErrorMsg(code), data, c)
}


// 分类下所有文章
func (a *ArticleApi)GetCateArt(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	Size,_ := strconv.Atoi(c.DefaultQuery("size","3"))
	Page,_ := strconv.Atoi(c.DefaultQuery("page","1"))
	data, code, total := model.GetCateArt(id, Size, Page)
	response.ResultAll(code, errmsg.GetErrorMsg(code), data, total, c)
}

// 文章编辑
func (a *ArticleApi)EditArticle(c *gin.Context)  {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code = model.UpdateArticle(id, &data)

	response.Result(code, errmsg.GetErrorMsg(code), "", c)
}

// 文章删除
func (a *ArticleApi)DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteArticle(id)

	response.Result(code, errmsg.GetErrorMsg(code), "", c)
}
