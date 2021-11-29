package model

import (
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Category	Category
	Title		string 	`gorm:"type:varchar(100);not null" json:"title"`
	Cid			int		`gorm:"type:int;not null" json:"cid"`
	Desc		string	`gorm:"type:varchar(200)" json:"desc""`
	Content		string	`gorm:"type:longtext" json:"content"`
	Img			string	`gorm:"type:varchar(100)" json:"img"`

}

//CreateArticle 新增文章
func CreateArticle(data *Article) int {
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//TODO: 查询分类下所有文章

//TODO: 查询单个文章

//TODO: GetArticles 查询文章列表，传pageSize,pageNum,返回User列表的切片,需要总数

func GetArticles(pageSize, pageNum int) []Category {
	var cate []Category
	err = db.Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&cate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cate
}

//EditArticle 编辑文章信息
func EditArticle(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err = db.Model(&art).Where("id = ?",id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//DeleteArticle 删除文章
func DeleteArticle(id int) int {
	var art Article
	err = db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}