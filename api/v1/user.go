package v1

import (
	"mygin/model"
	"mygin/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

var code int

//添加用户
func AddUser(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)
	code = model.CheckUser(user.Username)
	if code == errmsg.SUCCESS {
		code = model.CreateUser(&user)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    user,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个用户

//查询用户列表
func GetUsers(c *gin.Context) {

}

//编辑用户
func EditUser(c *gin.Context) {

}

//删除用户
func DeleteUser(c *gin.Context) {}
