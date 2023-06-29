package model

import "mygin/utils/errmsg"

type Category struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

func CheckCate(name string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS
}

func CreateCate(data *Category) int {
	result := db.Create(data)
	if result.Error != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetCate(pageSize, pageNum int) []Category {
	var cates []Category
	result := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates)
	if result.Error != nil {
		return nil
	}
	return cates
}

func DeleteCate(id int) int {
	var cate Category
	result := db.Where("id = ?", id).Delete(&cate)
	if result.Error != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func UpdateCate(id int, data *Category) int {
	var cate Category
	maps := make(map[string]interface{})
	maps["name"] = data.Name
	result := db.Model(&cate).Where("id = ?", id).Updates(maps)
	if result.Error != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// todo 查询分类下的所有文章
