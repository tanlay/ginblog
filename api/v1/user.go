package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"strconv"
)

/*
查询用户是否存在
添加用户
查询单个用户
查询用户列表
编辑用户
删除用户
 */

var code int

//AddUser 添加用户
func AddUser(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)
	code = model.CheckUser(user.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&user)
	} else if code == errmsg.ERROR_USER_EXIST {
		code = errmsg.ERROR_USER_EXIST
	}
	c.JSON(200, gin.H{
		"status": code,
		"data": user,
		"message": errmsg.GetErrMsg(code),
	})
}

//GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _:= strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetUsers(pageSize,pageNum)
	code = errmsg.SUCCESS
	c.JSON(200,gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
}

//EditUser 编辑用户
func EditUser(c *gin.Context) {

}

//DeleteUser 删除用户
func DeleteUser(c *gin.Context) {

}