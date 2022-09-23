package v1

import (
	"blog/model"
	"blog/model/response"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentApi struct {
}

// AddComment 新增评论
func (co *CommentApi) AddComment(c *gin.Context) {
	var data model.Comment
	_ = c.ShouldBindJSON(&data)
	code := data.AddComment()
	response.Result(code, data, c)
}

// GetComment 获取单个评论信息
func (co *CommentApi) GetComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := model.Comment{}
	code := data.GetComment(id)
	response.Result(code, data, c)
}

// DeleteComment 删除评论
func (co *CommentApi) DeleteComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteComment(uint(id))
	response.Result(code, "", c)
}

// GetCommentCount 获取评论数量
func (co *CommentApi) GetCommentCount(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	total := model.GetCommentCount(id)
	c.JSON(http.StatusOK, gin.H{
		"total": total,
	})
}

// GetCommentList 后台查询评论列表
func (co *CommentApi) GetCommentList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}
	data := model.Comments{}
	total, code := data.GetCommentList(pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrorMsg(code),
	})

}

// GetCommentListFront 展示页面显示评论列表
func (co *CommentApi) GetCommentListFront(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}
	data := model.Comments{}
	total, code := data.GetCommentListFront(id, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrorMsg(code),
	})

}

// CheckComment 通过审核
func (co *CommentApi) CheckComment(c *gin.Context) {
	var data model.Comment
	_ = c.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(c.Param("id"))

	code := data.CheckComment(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMsg(code),
	})
}

// UncheckComment 撤下评论审核
func (co *CommentApi) UncheckComment(c *gin.Context) {
	var data model.Comment
	_ = c.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(c.Param("id"))

	code := data.UncheckComment(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMsg(code),
	})
}
