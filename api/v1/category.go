package v1

import (
	"mygin/model"
	"mygin/utils/common"
	"mygin/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//查询分类是否存在

//添加分类

//查询单个分类下的文章

//查询分类列表

//编辑分类

//删除分类

//添加分类
func AddCate(c *gin.Context) {
	var cate model.Category
	err := c.ShouldBindJSON(&cate)
	if err != nil {
		code = errmsg.ERROR_VALIDATE_WRONG
	} else {
		code = model.CheckCate(cate.Name)
		if code == errmsg.SUCCESS {
			code = model.CreateCate(&cate)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    cate,
		"message": errmsg.GetErrMsg(code),
	})
}

// todo 查询分类下的所有文章

//查询分类列表
func GetCate(c *gin.Context) {
	var pageInfo common.PageInfo
	var cates []model.Category
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		code = errmsg.ERROR_QUERY_WRONG
		cates = nil
	} else {
		pageSize := pageInfo.PageSize
		pageNum := pageInfo.PageNum
		if pageSize == 0 {
			pageSize = -1
		}
		if pageNum == 0 {
			pageNum = 2
		}
		cates = model.GetCate(pageSize, pageNum)
		code = errmsg.SUCCESS
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    cates,
		"message": errmsg.GetErrMsg(code),
	})

}

//编辑分类
func EditCate(c *gin.Context) {
	var cate model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&cate)
	if err != nil {
		code = errmsg.ERROR_QUERY_WRONG
	} else {
		code = model.CheckCate(cate.Name)
		if code == errmsg.SUCCESS {
			code = model.UpdateCate(id, &cate)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除分类
func DeleteCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteCate(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
