package model

import (
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Category	Category	`gorm:"foreignkey:Cid"`
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

//GetCateArticle 查询分类下所有文章
func GetCateArticle(id, pageSize, pageNum int) ([]Article,int,int64) {
	var (
		cateArtList []Article
		total int64
	)
	result := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid =?",id).Find(&cateArtList)
	total = result.RowsAffected
	if result.Error != nil {
		return nil, errmsg.ERROR_CATEGORY_NOT_EXIST,0
	}
	return cateArtList, errmsg.SUCCESS,total
}

//GetArticleInfo 查询单个文章
func GetArticleInfo(id int) (Article,int){
	var art Article
	err = db.Preload("Category").Where("id= ?",id).First(&art).Error
	if err != nil  {
		return art, errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return art, errmsg.SUCCESS
}


//GetArticles 查询文章列表，传pageSize,pageNum,返回User列表的切片,需要总数
func GetArticles(pageSize, pageNum int) ([]Article,int,int64) {
	var (
		articleList []Article
		total	int64
	)
	result := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&articleList)
	total = result.RowsAffected
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR,0
	}
	return articleList,errmsg.SUCCESS,total
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