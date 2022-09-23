package model

import (
	"blog/global"
	"blog/model/schemas"
	"blog/utils/errmsg"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Avatar   string `gorm:"type:varchar(40);not null" json:"avatar" label:"头像"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required" label:"角色"`
}

type Users []User

// 存在性判断
func CheckUser(name string) (code int) {
	var users User
	// 查询用户是否存在
	global.Db.Select("id").Where("username=?", name).First(&users)
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
	global.Db.Select("id").Where("username=?", name).First(&user)
	if user.ID > 0 {
		if int(user.ID) == id {
			return errmsg.SUCCSE
		}
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCSE
}

// 添加用户
func (model *User) CreateUser() int {
	// 添加时接收一下错误
	err := global.Db.Create(&model).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 单个用户
func (user *User) GetUser(id int) int {
	err := global.Db.Select("username,avatar,role").Where("ID=?", id).First(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 用户列表
func (users *Users) GetUsers(username string, Size int, Page int) int64 {
	var total int64
	if username != "" {
		if username != "" {
			global.Db.Select("id,username,role").Where(
				"username LIKE ?", username+"%",
			).Limit(Size).Offset((Page - 1) * Size).Find(&users)
			global.Db.Model(&users).Where(
				"username LIKE ?", username+"%",
			).Count(&total)
			return total
		}
	}
	// 分页
	err := global.Db.Select("id,username,role").Offset((Page - 1) * Size).Limit(Size).Find(&users).Error
	global.Db.Model(&users).Count(&total)
	if err != nil {
		return 0
	}
	// 返回用户的列表
	return total
}

// 编辑用户
func (model *User) UpdateUser(id int) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = model.Username
	maps["role"] = model.Role
	err := global.Db.Model(&user).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除用户
func DeleteUser(id int) int {
	var user User
	err := global.Db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
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

// 登陆验证
func (user *User) CheckLogin(data schemas.Login) int {
	var PasswordErr error
	global.Db.Where("username=?", data.Username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if PasswordErr != nil {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCSE
}

// 管理员重置密码
func AdminEdit(id int, password string) int {
	var user User
	var maps = make(map[string]interface{})
	password = ScryptPw(password)
	maps["password"] = password
	err := global.Db.Model(&user).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func NosOne() {
	var user User
	global.Db.First(&user)
	if user.ID > 0 {
		fmt.Println("成功初始化")
		fmt.Println("后台地址: /admin")
		fmt.Println("初始化管理员账号: Salmon")
		fmt.Println("初始化管理员密码: 123456")
		return
	}
	maps := map[string]interface{}{
		"username": "Salmon",
		"password": "$2a$10$OW34MAkDaCWTVKLBSVej.OXV/O00kqb0FuUtQ3t/PkncwBJbm.ujK",
		"avatar":   "null",
		"role":     1,
	}
	err := global.Db.Model(&User{}).Create(&maps).Error
	if err != nil {
		fmt.Println("启动脚本错误")
		return
	}
	global.Db.Where("username=?", "Salmon").Update("role", 1)
	fmt.Println("初始化成功")
	fmt.Println("后台地址: /admin")
	fmt.Println("初始化管理员账号: Salmon")
	fmt.Println("初始化管理员密码: 123456")
	return
}

// CheckLoginFront 前台登录
func (user *User) CheckLoginFront(username string, password string) int {
	var PasswordErr error

	global.Db.Where("username = ?", username).First(&user)

	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if PasswordErr != nil {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	return errmsg.SUCCSE
}
