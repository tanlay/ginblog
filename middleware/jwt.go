package middleware

import (
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

//JwtKey 秘钥
var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username	string	`json:"username"`
	jwt.StandardClaims
}

var code int

//GenerateToken 生成token
func GenerateToken(username string) (string,int){
	expiredTime := time.Now().Add(10*time.Hour)
	SetClaims := MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
			Issuer: "ginblog",
		},
	}

	reqClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaims.SignedString(JwtKey)
	if err != nil{
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCESS
}

//ParseToken 验证token
func ParseToken(token string) (*MyClaims, int) {
	parseToken, _ := jwt.ParseWithClaims(token,&MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, _ := parseToken.Claims.(*MyClaims); parseToken.Valid {
		return key, errmsg.SUCCESS
	}
	return nil, errmsg.ERROR
}

//jwt中间件，用于控制验证

func JwtToken() gin.HandlerFunc{
	return func (c *gin.Context) {
		//固定写法
		tokenHeader := c.Request.Header.Get("Authorization")
		code := errmsg.SUCCESS
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			c.JSON(200, gin.H{
				"code": code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		//固定写法
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_ERONG
			c.JSON(200, gin.H{
				"code": code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		key, tCode := ParseToken(checkToken[1])
		if tCode == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(200, gin.H{
				"code": code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ERROR_TOKEN_TIMEOUT
			c.JSON(200, gin.H{
				"code": code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		c.Set("username", key.Username)
		c.Next()
	}
}












