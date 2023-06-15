package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Category Category `gorm:"foreignKey:Cid"`
	Title    string   `gorm:"type:varchar(100);not null" json:"title"`
	Content  string   `gorm:"type:longtext" json:"content"`
	Desc     string   `gorm:"type:varchar(200)" json:"desc"`
	Cid      int      `gorm:"type:int;not null" json:"cid"`
	Img      string   `gorm:"type:varchar(100)" json:"img"`
}
