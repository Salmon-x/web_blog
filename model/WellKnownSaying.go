package model

import (
	"blog/global"
	"blog/utils/errmsg"
	"gorm.io/gorm"
)

type WellKnownSaying struct {
	Id    uint   `gorm:"primary_key;auto_increment" json:"id"`
	Title string `gorm:"type:varchar(200);not null" json:"title"`
}

type WellKnownSayings []WellKnownSaying

// 存在性判断
func Checkwks(title string) (code int) {
	var wks WellKnownSaying
	// 查询用户是否存在
	global.Db.Select("id").Where("title=?", title).First(&wks)
	if wks.Id > 0 {
		// 如果存在则引出错误
		return errmsg.ERROR_WKS_USED
	}
	// 不存在则输入正确
	return errmsg.SUCCSE
}

// 添加分类
func (model *WellKnownSaying) Createwks() int {
	// 添加时接收一下错误
	err := global.Db.Create(&model).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 分类列表
func (wks *WellKnownSayings) GetWks() {
	// 分页
	err := global.Db.Find(&wks).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	// 返回分类列表
	return
}

func (cate *WellKnownSaying) GetWksInfo(id int) int {
	global.Db.Where("id = ?", id).First(&cate)
	return errmsg.SUCCSE
}

// 编辑分类
func (model *WellKnownSaying) UpdateWks(id int) int {
	var wks WellKnownSaying
	var maps = make(map[string]interface{})
	maps["title"] = model.Title
	err := global.Db.Model(&wks).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除分类
func DeleteWks(id int) int {
	var cate WellKnownSaying
	err := global.Db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
