package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"strconv"
)

/*
查询分类是否存在
添加分类
查询单个分类下的文章
查询分类列表
编辑分类
删除分类
*/

//AddCategory 添加分类
func AddCategory(c *gin.Context) {
	var cate model.Category
	_ = c.ShouldBindJSON(&cate)
	code = model.CheckCategory(cate.Name)
	if code == errmsg.SUCCESS {
		model.CreateCategory(&cate)
	} else if code == errmsg.ERROR_CATEGORY_EXIST {
		code = errmsg.ERROR_CATEGORY_EXIST
	}
	c.JSON(200, gin.H{
		"status": code,
		"data": cate,
		"message": errmsg.GetErrMsg(code),
	})
}

//GetCategories 查询分类列表
func GetCategories(c *gin.Context) {
	pageSize, _:= strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetCategories(pageSize,pageNum)
	code = errmsg.SUCCESS
	c.JSON(200,gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
}

//EditCategory 编辑分类信息
func EditCategory(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.EditCategory(id, &data)
	} else if code == errmsg.ERROR_CATEGORY_EXIST {
		c.Abort()
	}
	c.JSON(200, gin.H{
		"static": code,
		"message": errmsg.GetErrMsg(code),
	})
}

//DeleteCategory 删除分类
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteCategory(id)
	c.JSON(200, gin.H{
		"static": code,
		"message": errmsg.GetErrMsg(code),
	})
}