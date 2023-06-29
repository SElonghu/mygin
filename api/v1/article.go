package v1

import (
	"mygin/model"
	"mygin/utils/common"
	"mygin/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//添加文章

//查询单个文章

//查询文章列表

//编辑文章

//删除文章

func AddArt(c *gin.Context) {
	var article model.Article
	err := c.ShouldBindJSON(&article)
	if err != nil {
		code = errmsg.ERROR_VALIDATE_WRONG
	} else {
		code = model.CreateArt(&article)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    article,
		"message": errmsg.GetErrMsg(code),
	})
}

// todo 查询分类下的所有文章

//查询文章列表
func GetArt(c *gin.Context) {
	var pageInfo common.PageInfo
	var articleList []model.Article
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		code = errmsg.ERROR_QUERY_WRONG
		articleList = nil
	} else {
		pageSize := pageInfo.PageSize
		pageNum := pageInfo.PageNum
		if pageSize == 0 {
			pageSize = -1
		}
		if pageNum == 0 {
			pageNum = 2
		}
		articleList, code = model.GetArt(pageSize, pageNum)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    articleList,
		"message": errmsg.GetErrMsg(code),
	})

}

//编辑文章
func EditArt(c *gin.Context) {
	var article model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&article)
	if err != nil {
		code = errmsg.ERROR_QUERY_WRONG
	} else {
		code = model.UpdateArt(id, &article)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除文章
func DeleteArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteArt(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
