package model

import (
	"mygin/utils/errmsg"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" binding:"required"`
	Password string `gorm:"type:varchar(20);not null" json:"password" binding:"required"`
	Role     int    `gorm:"type:int" json:"role"`
}

func CheckUser(name string) (code int) {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

func CreateUser(data *User) int {
	result := db.Create(data)
	if result.Error != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetUsers(pageSize, pageNum int) []User {
	var users []User
	result := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
	if result.Error != nil {
		return nil
	}
	return users
}
