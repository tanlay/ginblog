package model

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
)

//User 用户结构
type User struct {
	gorm.Model
	Username 	string	`gorm:"type:varchar(20);not null" json:"username"`
	Password	string	`gorm:"type:varchar(20);not null" json:"password"`
	Role		int	`gorm:"type:int" json:"role"`
}

//CheckUser 查询用户是否存在
func CheckUser(name string) (code int) {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USER_EXIST
	}
	return errmsg.SUCCESS
}

//CreateUser 新增用户
func CreateUser(data *User) int {
	data.Password = ScryptPwd(data.Password)
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetUsers 查询用户列表，传pageSize,pageNum,返回User列表的切片
func GetUsers(pageSize, pageNum int) []User {
	var users []User
	err = db.Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

//EditUser 编辑用户


//DeleteUser 删除用户


//ScryptPwd 密码加密,入参：password,返回密码加密后的字符串
func ScryptPwd(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12,28,45,76,35,11,12,35}
	HashPwd, err := scrypt.Key([]byte(password), salt,16384,8,1,KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fPwd := base64.StdEncoding.EncodeToString(HashPwd)
	return fPwd
}