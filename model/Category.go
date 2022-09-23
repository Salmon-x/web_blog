package model

import (
	"blog/global"
	"blog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

type Categorys []Category

// 存在性判断
func CheckCategory(name string) (code int) {
	var cate Category
	// 查询用户是否存在
	global.Db.Select("id").Where("name=?", name).First(&cate)
	if cate.ID > 0 {
		// 如果存在则引出错误
		return errmsg.ERROR_CATENAME_USED
	}
	// 不存在则输入正确
	return errmsg.SUCCSE
}

// 添加分类
func (model *Category) CreateCategory() int {
	// 添加时接收一下错误
	err := global.Db.Create(&model).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// GetCateInfo 查询单个分类信息
func (cate *Category) GetCateInfo(id int) int {
	global.Db.Where("id = ?", id).First(&cate)
	return errmsg.SUCCSE
}

// 分类列表
func (cates *Categorys) GetCategorys(Size int, Page int) int64 {
	var total int64
	// 分页
	err := global.Db.Limit(Size).Offset((Page - 1) * Size).Find(&cates).Error
	global.Db.Model(&cates).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0
	}
	// 返回分类列表
	return total
}

// 编辑分类
func (model *Category) UpdateCategory(id int) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = model.Name
	err := global.Db.Model(&cate).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除分类
func DeleteCategory(id int) int {
	var cate Category
	err := global.Db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
