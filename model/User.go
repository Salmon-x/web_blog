package model

import (
	"blog/utils/errmsg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
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

// 唯一性判断，在编辑用户的时候，如果使用之前的存在性判断，那只修改其他字段，而不修改username，会永远被卡在那条判断，为了解藕，重写一个唯一判断
func UniqueUser(name string, id int) int {
	var user User
	db.Select("id").Where("username=?", name).First(&user)
	if user.ID > 0{
		if int(user.ID) == id {
			return errmsg.SUCCSE
		}
		return errmsg.ERROR_USERNAME_USED
	}
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
func UpdateUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&user).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}


// 密码加密
func ScryptPw(password string) string {
	const cost = 10
	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}
	return string(HashPw)
}

// BeforeCreate 密码加密&权限控制
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	u.Role = 2
	return nil
}