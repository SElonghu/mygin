package model

import (
	"encoding/base64"
	"log"
	"mygin/utils/errmsg"

	"golang.org/x/crypto/scrypt"
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
	data.Password = CryptPW(data.Password)
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

func DeleteUser(id int) int {
	var user User
	result := db.Where("id = ?", id).Delete(&user)
	if result.Error != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func UpdateUser(id int, data *User) int {
	var user User
	maps := make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	result := db.Model(&user).Where("id = ?", id).Updates(maps)
	if result.Error != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func CryptPW(password string) string {
	salt := []byte{14, 54, 32, 8, 31, 6, 9, 25}
	dk, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, 10)
	if err != nil {
		log.Fatal(err)
	}
	finalPW := base64.StdEncoding.EncodeToString(dk)
	return finalPW
}
