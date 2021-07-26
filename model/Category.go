package model

import (
	"blog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// 存在性判断
func CheckCategory(name string)(code int) {
	var cate Category
	// 查询用户是否存在
	Db.Select("id").Where("name=?",name).First(&cate)
	if cate.ID > 0 {
		// 如果存在则引出错误
		return errmsg.ERROR_CATENAME_USED
	}
	// 不存在则输入正确
	return errmsg.SUCCSE
}

// 添加分类
func CreateCategory(data *Category)int  {
	// 添加时接收一下错误
	err := Db.Create(&data).Error
	if err!=nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// GetCateInfo 查询单个分类信息
func GetCateInfo(id int) (Category,int) {
	var cate Category
	Db.Where("id = ?",id).First(&cate)
	return cate,errmsg.SUCCSE
}

// 分类列表
func GetCategorys(Size int, Page int)([]Category,int64)  {
	var cates []Category
	var total int64
	// 分页
	err = Db.Limit(Size).Offset((Page - 1) * Size).Find(&cates).Error
	Db.Model(&cates).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil,0
	}
	// 返回分类列表
	return cates,total
}

// 编辑分类
func UpdateCategory(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = Db.Model(&cate).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除分类
func DeleteCategory(id int) int {
	var cate Category
	err = Db.Where("id = ?", id).Delete(&cate).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}