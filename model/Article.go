package model

import (
	"blog/utils/errmsg"
	"gorm.io/gorm"
)


type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title        string `gorm:"type:varchar(100);not null" json:"title"`
	Cid          int    `gorm:"type:int;not null" json:"cid"`
	Desc         string `gorm:"type:varchar(200)" json:"desc"`
	Content      string `gorm:"type:longtext" json:"content"`
	Img          string `gorm:"type:varchar(100)" json:"img"`
}


// 添加用户
func CreateArticle(data *Article)int  {
	// 添加时接收一下错误
	err := db.Create(&data).Error
	if err!=nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}


// 分类列表
func GetArticles(PageSize int, PageNum int)[]Article {
	var articles []Article
	// 分页
	err = db.Limit(PageSize).Offset((PageNum - 1) * PageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	// 返回分类列表
	return articles
}

// 编辑分类
func UpdateArticle(id int, data *Article) int {
	var article Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["cid"] = data.Cid
	maps["img"] = data.Img
	err = db.Model(&article).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除分类
func DeleteArticle(id int) int {
	var article Article
	err = db.Where("id = ?", id).Delete(&article).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}