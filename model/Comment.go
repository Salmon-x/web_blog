package model

import (
	"blog/global"
	"blog/utils/errmsg"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserId    uint   `json:"user_id"`
	ArticleId uint   `json:"article_id"`
	Title     string `json:"article_title"`
	Username  string `json:"username"`
	Content   string `gorm:"type:varchar(500);not null;" json:"content"`
	Status    int8   `gorm:"type:tinyint;default:2" json:"status"`
}

type Comments []Comment

// AddComment 新增评论
func (data *Comment) AddComment() int {
	err := global.Db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// GetComment 查询单个评论
func (comment *Comment) GetComment(id int) int {
	err := global.Db.Where("id = ?", id).First(&comment).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// GetCommentList 后台所有获取评论列表
func (comments *Comments) GetCommentList(pageSize int, pageNum int) (int64, int) {
	var total int64
	global.Db.Find(&comments).Count(&total)
	err := global.Db.Model(&comments).Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Select("comment.id, article.title,user_id,article_id, user.username, comment.content, comment.status,comment.created_at,comment.deleted_at").Joins("LEFT JOIN article ON comment.article_id = article.id").Joins("LEFT JOIN user ON comment.user_id = user.id").Scan(&comments).Error
	if err != nil {
		return 0, errmsg.ERROR
	}
	return total, errmsg.SUCCSE
}

// GetCommentCount 获取评论数量
func GetCommentCount(id int) int64 {
	var comment Comment
	var total int64
	global.Db.Find(&comment).Where("article_id = ?", id).Where("status = ?", 1).Count(&total)
	return total
}

// GetCommentListFront 展示页面获取评论列表
func (comments *Comments) GetCommentListFront(id int, pageSize int, pageNum int) (int64, int) {
	var total int64
	global.Db.Find(&Comment{}).Where("article_id = ?", id).Where("status = ?", 1).Count(&total)
	err := global.Db.Model(&Comment{}).Limit(pageSize).Offset((pageNum-1)*pageSize).Order("Created_At DESC").Select("comment.id, article.title, user_id, article_id, user.username, comment.content, comment.status,comment.created_at,comment.deleted_at").Joins("LEFT JOIN article ON comment.article_id = article.id").Joins("LEFT JOIN user ON comment.user_id = user.id").Where("article_id = ?",
		id).Where("status = ?", 1).Scan(&comments).Error
	if err != nil {
		return 0, errmsg.ERROR
	}
	return total, errmsg.SUCCSE
}

// 编辑评论（暂不允许编辑评论）

// DeleteComment 删除评论
func DeleteComment(id uint) int {
	var comment Comment
	err := global.Db.Where("id = ?", id).Delete(&comment).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// CheckComment 通过评论
func (data *Comment) CheckComment(id int) int {
	var comment Comment
	var res Comment
	var article Article
	var maps = make(map[string]interface{})
	maps["status"] = data.Status

	err := global.Db.Model(&comment).Where("id = ?", id).Updates(maps).First(&res).Error
	global.Db.Model(&article).Where("id = ?", res.ArticleId).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1))
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// UncheckComment 撤下评论
func (data *Comment) UncheckComment(id int) int {
	var comment Comment
	var res Comment
	var article Article
	var maps = make(map[string]interface{})
	maps["status"] = data.Status

	err := global.Db.Model(&comment).Where("id = ?", id).Updates(maps).First(&res).Error
	global.Db.Model(&article).Where("id = ?", res.ArticleId).UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1))
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
