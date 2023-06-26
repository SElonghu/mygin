package v1

import (
	"mygin/model"
	"mygin/utils/common"
	"mygin/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

var code int

//添加用户
func AddUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		code = errmsg.ERROR_VALIDATE_WRONG
	} else {
		code = model.CheckUser(user.Username)
		if code == errmsg.SUCCESS {
			code = model.CreateUser(&user)
		}
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
	var pageInfo common.PageInfo
	var users []model.User
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		code = errmsg.ERROR_QUERY_WRONG
		users = nil
	} else {
		pageSize := pageInfo.PageSize
		pageNum := pageInfo.PageNum
		if pageSize == 0 {
			pageSize = -1
		}
		if pageNum == 0 {
			pageNum = 2
		}
		users = model.GetUsers(pageSize, pageNum)
		code = errmsg.SUCCESS
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    users,
		"message": errmsg.GetErrMsg(code),
	})

}

//编辑用户
func EditUser(c *gin.Context) {

}

//删除用户
func DeleteUser(c *gin.Context) {}
