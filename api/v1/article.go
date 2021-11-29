package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"strconv"
)

/*
创建文章
查询单个文章
查询文章列表
编辑文章
删除文章
*/

//AddArticle 添加文章
func AddArticle(c *gin.Context) {
	var art model.Article
	_ = c.ShouldBindJSON(&art)

	code = model.CreateArticle(&art)

	c.JSON(200, gin.H{
		"status": code,
		"data": art,
		"message": errmsg.GetErrMsg(code),
	})
}

//GetCateArticle 查询分类下所有文章信息
func GetCateArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pageSize, _:= strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code := model.GetCateArticle(id, pageSize, pageNum)
	c.JSON(200,gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
}

//GetArticleInfo 查询单个文章
func GetArticleInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetArticleInfo(id)
	c.JSON(200,gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
}

//TODO:GetArticles 查询文章列表

func GetArticles(c *gin.Context) {
	pageSize, _:= strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code := model.GetArticles(pageSize,pageNum)
	c.JSON(200,gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
}

//EditArticle 编辑文章信息
func EditArticle(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)

	code = model.EditArticle(id, &data)

	c.JSON(200, gin.H{
		"static": code,
		"message": errmsg.GetErrMsg(code),
	})
}

//DeleteArticle 删除分类
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteArticle(id)
	c.JSON(200, gin.H{
		"static": code,
		"message": errmsg.GetErrMsg(code),
	})
}