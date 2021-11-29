package v1

import (
	"ginblog/middleware"
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var (
		data model.User
		token string
		code int
	)
	c.ShouldBindJSON(&data)
	code = model.CheckLogin(data.Username,data.Password)
	if code == errmsg.SUCCESS{
		token, code = middleware.GenerateToken(data.Username)
	}
	c.JSON(200, gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
		"token": token,
	})
}
