package model

import (
	"mygin/utils/errmsg"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Category Category `gorm:"foreignKey:Cid"`
	Title    string   `gorm:"type:varchar(100);not null" json:"title"`
	Content  string   `gorm:"type:longtext" json:"content"`
	Desc     string   `gorm:"type:varchar(200)" json:"desc"`
	Cid      int      `gorm:"type:int;not null" json:"cid"`
	Img      string   `gorm:"type:varchar(100)" json:"img"`
}

func CreateArt(data *Article) int {
	result := db.Create(data)
	if result.Error != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// todo 查询单个文章

func GetArt(pageSize, pageNum int) ([]Article,int) {
	var articles []Article
	result := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles)
	if result.Error != nil {
		return nil, errmsg.ERROR
	}
	return articles, errmsg.SUCCESS
}

func DeleteArt(id int) int {
	var article Article
	result := db.Where("id = ?", id).Delete(&article)
	if result.Error != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func UpdateArt(id int, data *Article) int {
	var article Article
	maps := make(map[string]interface{})
	maps["title"] = data.Title
	maps["content"] = data.Content
	maps["desc"] = data.Desc
	maps["cid"] = data.Cid
	maps["img"] = data.Img
	result := db.Model(&article).Where("id = ?", id).Updates(maps)
	if result.Error != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
