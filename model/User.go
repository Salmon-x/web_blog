package model

import (
	"blog/utils/errmsg"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Avatar string `gorm:"type:varchar(40);not null" json:"avatar" label:"头像"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色"`
}

// 存在性判断
func CheckUser(name string)(code int) {
	var users User
	// 查询用户是否存在
	db.Select("id").Where("username=?",name).First(&users)
	if users.ID > 0 {
		// 如果存在则引出错误
		return errmsg.ERROR_USERNAME_USED
	}
	// 不存在则输入正确
	return errmsg.SUCCSE
}

// 添加用户
func CreateUser(data *User)int  {
	// 添加时接收一下错误
	err := db.Create(&data).Error
	if err!=nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 用户列表
func GetUsers(PageSize int, PageNum int)[]User  {
	var users []User
	// 分页
	err = db.Limit(PageSize).Offset((PageNum - 1) * PageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	// 返回用户的列表
	return users
}

// 编辑用户
func UpdateUser()  {
	
}