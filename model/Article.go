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


// 添加文章
func CreateArticle(data *Article)int  {
	// 添加时接收一下错误
	err := Db.Create(&data).Error
	if err!=nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 分类下所有文章
func GetCateArt(cid int, Size int, Page int)([]Article, int, int64){
	var cateArtList []Article
	var total int64
	err = Db.Order("Updated_At DESC").Preload("Category").Limit(Size).Offset((Page - 1) * Size).Where("cid=?", cid).Find(&cateArtList).Error
	Db.Model(&cateArtList).Where("cid=?", cid).Count(&total)
	if err != nil{
		return nil,errmsg.ERROR_CATE_NOT_EXIST,0
	}
	return cateArtList, errmsg.SUCCSE, total
}

// SearchArticle 搜索文章标题
func SearchArticle(title string, pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var err error
	var total int64
	err = Db.Order("Updated_At DESC").Select("article.id,title, img, created_at, updated_at, `desc`, comment_count, read_count, category.name").Limit(pageSize).Offset((pageNum-1)*pageSize).Order("Created_At DESC").Joins("Category").Where("title LIKE ?",
		title+"%",
	).Find(&articleList).Error
	// 单独计数
	Db.Model(&articleList).Count(&total)

	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCSE, total
}


// 单个文章
func ArticleInfo(id int) (Article, int) {
	var art Article
	err = Db.Preload("Category").Where("id=?",id).First(&art).Error
	if err != nil{
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCSE
}

// 文章列表
func GetArticles(Size int, Page int)([]Article, int, int64) {
	var articles []Article
	var total int64
	// 分页
	err = Db.Order("Updated_At DESC").Preload("Category").Limit(Size).Offset((Page - 1) * Size).Find(&articles).Error
	Db.Model(&articles).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR,0
	}
	// 返回文章列表
	return articles, errmsg.SUCCSE,total
}

// 编辑文章
func UpdateArticle(id int, data *Article) int {
	var article Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["cid"] = data.Cid
	maps["img"] = data.Img
	err = Db.Model(&article).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除文章
func DeleteArticle(id int) int {
	var article Article
	err = Db.Where("id = ?", id).Delete(&article).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}