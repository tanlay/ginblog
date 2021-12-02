package model

import (
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Category struct {
	ID		uint	`gorm:"primary_key;auto_increment" json:"id"`
	Name	string	`gorm:"type:varchar(20);not null" json:"name"`
}


//CheckCategory 查询分类是否存在
func CheckCategory(name string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATEGORY_EXIST
	}
	return errmsg.SUCCESS
}

//CreateCategory 新增分类
func CreateCategory(data *Category) int {
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询分类下的所有文章


// GetCategories 查询分类列表，传pageSize,pageNum,返回User列表的切片
func GetCategories(pageSize, pageNum int) ([]Category,int64) {
	var (
		cate []Category
		total int64
	)
	result := db.Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&cate)
	total = result.RowsAffected
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cate, total
}

//EditCategory 编辑分类信息
func EditCategory(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&cate).Where("id = ?",id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//DeleteCategory 删除分类
func DeleteCategory(id int) int {
	var cate Category
	err = db.Where("id = ?",id).Delete(&cate).Error
	if err !=  nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}